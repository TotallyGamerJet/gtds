// +build darwin

package gtds

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

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
    @autoreleasepool {

    for (;;) {
        NSEvent* event = [NSApp nextEventMatchingMask:NSEventMaskAny
			untilDate:[NSDate distantPast]
			   inMode:NSDefaultRunLoopMode
			  dequeue:YES];
        if (event == nil)
            break;

        [NSApp sendEvent:event];
    }

    } // autoreleasepool
}

void* CreateWindow(_GoString_ title, int width, int height, int style) {
	[NSAutoreleasePool new];
	id window = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, width, height)
		styleMask:style backing:NSBackingStoreBuffered defer:NO];
	[window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
	//[window setCanBecomeKeyWindow:YES];
	[window setTitle:[[[NSString alloc] initWithBytes:_GoStringPtr(title) length:_GoStringLen(title) encoding:NSUTF8StringEncoding]autorelease]];
	[window makeKeyAndOrderFront:nil];
	[window setDelegate:[[WindowDelegate alloc] init]];

	platformPollEvents(); //Allow NSApp to catch up
	return window;
}

NSRect platformGetFrameBufferSize(void* window) {
	const NSRect contentRect = [((NSWindow*)(window)).contentView frame];
    const NSRect fbRect = [((NSWindow*)(window)).contentView convertRectToBacking:contentRect];
	return fbRect;
}

void * Window_ContentView(void * window) {
	return ((NSWindow *)window).contentView;
}

void View_SetLayer(void * view, void * layer) {
	((NSView *)view).layer = (CALayer *)layer;
}

void View_SetWantsLayer(void * view, BOOL wantsLayer) {
	((NSView *)view).wantsLayer = wantsLayer;
}
*/
import "C"
import (
	"letsgo/internal/coreanim"
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
	win := Window{ptr}
	getData(win)
	return win
}

func platformInit() {
	C.StartApp()
}

func platformPollEvents() {
	C.platformPollEvents()
}

// View is the infrastructure for drawing, printing, and handling events in an app.
//
// Reference: https://developer.apple.com/documentation/appkit/nsview.
type View struct {
	view unsafe.Pointer
}

// SetLayer sets v.layer to l.
//
// Reference: https://developer.apple.com/documentation/appkit/nsview/1483298-layer.
func (v View) SetLayer(l coreanim.Layer) {
	C.View_SetLayer(v.view, l.Layer())
}

// SetWantsLayer sets v.wantsLayer to wantsLayer.
//
// Reference: https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer.
func (v View) SetWantsLayer(wantsLayer bool) {
	C.View_SetWantsLayer(v.view, toCBool(wantsLayer))
}

func toCBool(b bool) C.BOOL {
	if b {
		return 1
	}
	return 0
}

func (w Window) ContentView() View {
	return View{C.Window_ContentView(w.ptr)}
}

func (w Window) GetFrameBufferSize() (int, int) {
	rect := C.platformGetFrameBufferSize(w.ptr)
	return int(rect.size.width), int(rect.size.height)
}
