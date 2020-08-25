// +build darwin

package gtds

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

@interface AppDelegate : NSObject <NSApplicationDelegate>
@property (assign) NSWindow *window;
@end

@implementation AppDelegate
@synthesize window;
- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
    // Insert code here to initialize your application
    [NSApp activateIgnoringOtherApps:YES];
}
- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)sender {
	return YES;
}
@end

@interface WindowDelegate : NSObject <NSWindowDelegate>
@end

@implementation WindowDelegate
- (BOOL)windowShouldClose:(NSWindow *)sender {
	NSLog(@"closing");
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

void CreateWindow(_GoString_ title, int width, int height, int style) {
	[NSAutoreleasePool new];
	id window = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, width, height)
		styleMask:style backing:NSBackingStoreBuffered defer:NO];
	[window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
	//[window setCanBecomeKeyWindow:YES];
	[window setTitle:[[[NSString alloc] initWithBytes:_GoStringPtr(title) length:_GoStringLen(title) encoding:NSUTF8StringEncoding]autorelease]];
	[window makeKeyAndOrderFront:nil];
	[window setDelegate:[[WindowDelegate alloc] init]];
}
*/
import "C"
import (
	"github.com/faiface/mainthread"
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
	if style&Minimizable != 0 {
		nsStyle |= NSWindowStyleMaskMiniaturizable
	}
	return nsStyle
}

func platformCreateWindow(w WindowConfig) Window {
	mainthread.Call(func() {
		C.CreateWindow(w.Title, C.int(w.Width), C.int(w.Height), C.int(translateStyle(w.Style)))
	})
	return Window{}
}

func platformRun() {
	C.StartApp()
}
