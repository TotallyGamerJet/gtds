// +build darwin

package gtds

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

extern void appUpdate();

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
- (void)applicationWillUpdate:(NSNotification *)notification {
	appUpdate();
}
@end

@interface WindowDelegate : NSObject <NSWindowDelegate>
@end

@implementation WindowDelegate
- (BOOL)windowShouldClose:(NSWindow *)sender {
	NSLog(@"closing");
	return YES;
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

void CreateWindow(_GoString_ title, int width, int height) {
	[NSAutoreleasePool new];
	id window = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, width, height)
		styleMask:NSWindowStyleMaskTitled|NSWindowStyleMaskClosable backing:NSBackingStoreBuffered defer:NO];
	[window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
	[window setTitle:[[[NSString alloc] initWithBytes:_GoStringPtr(title) length:_GoStringLen(title) encoding:NSUTF8StringEncoding]autorelease]];
	[window makeKeyAndOrderFront:nil];
	[window setDelegate:[[WindowDelegate alloc] init]];
}
*/
import "C"

func platformCreateWindow(w WindowConfig) {
	C.CreateWindow(w.Title, C.int(w.Width), C.int(w.Height))
}

func platformRun() {
	C.StartApp()
}
