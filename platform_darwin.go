// +build darwin

package gtds

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework MetalKit -framework Metal

#import <Cocoa/Cocoa.h>
#import <MetalKit/MetalKit.h>
#import <Metal/Metal.h>

extern void shouldClose(void* window);

@interface AppDelegate : NSObject <NSApplicationDelegate>
@property (assign) NSWindow *window;
@end

@implementation AppDelegate
@synthesize window;
- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
    // Insert code here to initialize your application
    [NSApp activateIgnoringOtherApps:YES];
	[NSApp stop:nil];
}
- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)sender {
	return YES;
}
- (void)applicationWillUpdate:(NSNotification *)notification {

}
@end

@interface WindowDelegate : NSObject <NSWindowDelegate>
@end

@implementation WindowDelegate
- (BOOL)windowShouldClose:(NSWindow *)sender {
	shouldClose(sender);
	return YES; //TODO: check if other windows are open
}
@end

void StartApp() {
    [NSAutoreleasePool new];
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

	NSMenu* mainMenu = [[[NSMenu alloc] initWithTitle:@""] autorelease];
	[NSApp setMainMenu:mainMenu];
	// Application menu
	NSMenuItem* appleItem = [mainMenu addItemWithTitle:@""
		action:nil keyEquivalent:@""];
	NSString* appName = [[NSProcessInfo processInfo] processName];
	NSMenu* appleMenu = [[NSMenu alloc] initWithTitle:@""];
	// Apple menu
	[appleMenu addItemWithTitle:[@"About " stringByAppendingString:appName]
		action:@selector(orderFrontStandardAboutPanel:) keyEquivalent:@""];
	// Quit
	[appleMenu addItemWithTitle:[@"Quit " stringByAppendingString:appName]
		action:@selector(terminate:) keyEquivalent:@"q"];
	[appleItem setSubmenu:[appleMenu autorelease]];

	AppDelegate* delegate = [AppDelegate new];
	[NSApp setDelegate:delegate];
    [NSApp run];
}

void platformPollEvents(void) {
   NSAutoreleasePool* pool = [[NSAutoreleasePool alloc] init];
    for (;;) {
        NSEvent* event = [NSApp nextEventMatchingMask:NSEventMaskAny
			untilDate:[NSDate distantPast]
			   inMode:NSDefaultRunLoopMode
			  dequeue:YES];
        if (event == nil)
            break;

        [NSApp sendEvent:event];
    }
	[pool release];
}

NSRect platformGetFrameBufferSize(void* window) {
	const NSRect contentRect = [((NSWindow*)(window)).contentView frame];
    const NSRect fbRect = [((NSWindow*)(window)).contentView convertRectToBacking:contentRect];
	return fbRect;
}

void* CreateWindow(_GoString_ title, int width, int height, int style) {
	[NSAutoreleasePool new];
	id window = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, width, height)
		styleMask:style backing:NSBackingStoreBuffered defer:NO];
	[window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
	//[window setCanBecomeKeyWindow:YES];
	[window setTitle:[[[NSString alloc] initWithBytes:_GoStringPtr(title) length:_GoStringLen(title) encoding:NSUTF8StringEncoding]autorelease]];
	[window makeKeyAndOrderFront:nil];
	[window setDelegate:[[WindowDelegate alloc] init]]; //TODO: pass in the device
	MTKView* view = [[MTKView alloc]initWithFrame: [window frame] device: MTLCreateSystemDefaultDevice()];
	view.drawableSize = (CGSize)platformGetFrameBufferSize(window).size;
	view.paused = NO;
	view.enableSetNeedsDisplay = NO;
	((NSWindow *)window).contentView=view;
	platformPollEvents(); //Allow NSApp to catch up
	return window;
}

uint16_t Window_PixelFormat(void* window) {
	id view = ((NSWindow*)window).contentView;
	return ((MTKView*)view).colorPixelFormat;
}

void* Window_nextDrawable(void* window) {
	id layer = ((NSWindow*)window).contentView.layer;
	return [(CAMetalLayer *)layer nextDrawable];
}
*/
import "C"
import (
	"letsgo/internal/coreanim"
	mtl "letsgo/internal/metal"
	"unsafe"
)

func translateStyle(style WindowStyle) int {
	const (
		NSWindowStyleMaskBorderless             = 0
		NSWindowStyleMaskTitled                 = 1 << 0
		NSWindowStyleMaskClosable               = 1 << 1
		NSWindowStyleMaskMiniaturizable         = 1 << 2
		NSWindowStyleMaskResizable              = 1 << 3
		NSWindowStyleMaskUtilityWindow          = 1 << 4
		NSWindowStyleMaskDocModalWindow         = 1 << 6
		NSWindowStyleMaskUnifiedTitleAndToolbar = 1 << 12
		NSWindowStyleMaskHUDWindow              = 1 << 13
		NSWindowStyleMaskFullScreen             = 1 << 14
		NSWindowStyleMaskFullSizeContentView    = 1 << 15
	)
	if style == Borderless {
		return NSWindowStyleMaskBorderless
	}
	var nsStyle int
	if style&Titled != 0 {
		nsStyle |= NSWindowStyleMaskTitled
	}
	if style&Closable != 0 {
		nsStyle |= NSWindowStyleMaskClosable
	}
	if style&Resizable != 0 {
		nsStyle |= NSWindowStyleMaskResizable
	}
	if style&Hideable != 0 {
		nsStyle |= NSWindowStyleMaskMiniaturizable
	}
	if style&Fullscreen != 0 {
		nsStyle |= NSWindowStyleMaskFullScreen | NSWindowStyleMaskFullSizeContentView
	}
	return nsStyle
}

func platformCreateWindow(w WindowConfig) Window {
	ptr := C.CreateWindow(w.Title, C.int(w.Width), C.int(w.Height), C.int(translateStyle(w.Style)))
	win := Window{uintptr(ptr)}
	getData(win)
	return win
}

func platformInit() {
	C.StartApp()
}

func platformPollEvents() {
	C.platformPollEvents()
}

func (w Window) GetFrameBufferSize() (int, int) {
	rect := C.platformGetFrameBufferSize(unsafe.Pointer(w.ptr))
	return int(rect.size.width), int(rect.size.height)
}

func (w Window) PixelFormat() mtl.PixelFormat {
	return mtl.PixelFormat(C.Window_PixelFormat(unsafe.Pointer(w.ptr)))
}

func (w Window) NextDrawable() (coreanim.MetalDrawable, error) {
	return coreanim.MetalDrawable{C.Window_nextDrawable(unsafe.Pointer(w.ptr))}, nil
}
