package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Cocoa/Cocoa.h>

@interface AppDelegate : NSObject <NSApplicationDelegate>
@end

@implementation AppDelegate

- (void)applicationDidChangeScreenParameters:(NSNotification *) notification {}

- (void)applicationWillFinishLaunching:(NSNotification *)notification {
	// In case we are unbundled, make us a proper UI application
	[NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

	// Menu bar setup must go between sharedApplication and finishLaunching
	// in order to properly emulate the behavior of NSApplicationMain
//         createMenuBar();
}

- (void)applicationDidFinishLaunching:(NSNotification *)notification {
	NSLog(@"app started");
    //_glfwPlatformPostEmptyEvent();
    [NSApp stop:nil];
}

- (void)applicationDidHide:(NSNotification *)notification {}

@end

void _platformInit() {
	@autoreleasepool {
		[NSApplication sharedApplication];
		[NSApp setDelegate:[[AppDelegate alloc] init]];
		if (![[NSRunningApplication currentApplication] isFinishedLaunching])
        	[NSApp run];
	}
}

@interface WindowDelegate : NSObject
@end

@implementation WindowDelegate

- (BOOL)windowShouldClose:(id)sender {
    NSLog(@"shouldClose");
    return NO;
}

- (void)windowDidResize:(NSNotification *)notification {
    NSLog(@"didResize");
}

- (void)windowDidMove:(NSNotification *)notification {
    NSLog(@"didMove");
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
   NSLog(@"didMiniaturize");
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
    NSLog(@"didDeminiaturize");
}

- (void)windowDidBecomeKey:(NSNotification *)notification {
    NSLog(@"didBecomeKey");
}

- (void)windowDidResignKey:(NSNotification *)notification {
    NSLog(@"didResignKey");
}

- (void)windowDidChangeOcclusionState:(NSNotification* )notification {
    NSLog(@"occlusion state");
}

@end

@interface Window : NSWindow {}
@end

@implementation Window

- (BOOL)canBecomeKeyWindow {
    // Required for NSWindowStyleMaskBorderless windows
    return YES;
}

- (BOOL)canBecomeMainWindow {
    return YES;
}

@end

@interface ContentView : NSView
@end

@implementation ContentView

- (BOOL)isOpaque{
    return YES;
}

- (BOOL)canBecomeKeyView {
    return YES;
}

- (BOOL)acceptsFirstResponder {
    return YES;
}

- (BOOL)wantsUpdateLayer {
    return YES;
}

- (void)updateLayer {
    NSLog(@"updateLayer");
}

- (void)cursorUpdate:(NSEvent *)event {
   NSLog(@"cursorUpdate");
}

- (BOOL)acceptsFirstMouse:(NSEvent *)event {
    return YES;
}

- (void)mouseDown:(NSEvent *)event {
    NSLog(@"mouseDown");
}

- (void)mouseDragged:(NSEvent *)event {
    [self mouseMoved:event];
}

- (void)mouseUp:(NSEvent *)event {
     NSLog(@"mouseup");
}

- (void)mouseMoved:(NSEvent *)event {
     NSLog(@"mouseMoved");
}

- (void)rightMouseDown:(NSEvent *)event {
     NSLog(@"rightMouseDown");
}

- (void)rightMouseDragged:(NSEvent *)event {
    [self mouseMoved:event];
}

- (void)rightMouseUp:(NSEvent *)event {
     NSLog(@"rightMouseUp");
}

- (void)otherMouseDown:(NSEvent *)event {
    NSLog(@"otherMouseDown");
}

- (void)otherMouseDragged:(NSEvent *)event {
    [self mouseMoved:event];
}

- (void)otherMouseUp:(NSEvent *)event {
     NSLog(@"Othermouseup");
}

- (void)mouseExited:(NSEvent *)event {
    NSLog(@"mouseExited");
}

- (void)mouseEntered:(NSEvent *)event {
     NSLog(@"mouseEntered");
}

- (void)viewDidChangeBackingProperties {
     NSLog(@"viewDIdChangeBackingProps");
}

- (void)drawRect:(NSRect)rect {
   NSLog(@"drawRect");
}

- (void)updateTrackingAreas {
    NSLog(@"updateTrackingAreas");
}

- (void)keyDown:(NSEvent *)event {
    NSLog(@"keyDown");

    [self interpretKeyEvents:@[event]];
}

- (void)flagsChanged:(NSEvent *)event {
    NSLog(@"flagsChanged");
}

- (void)keyUp:(NSEvent *)event {
    NSLog(@"KeyUp");
}

- (void)scrollWheel:(NSEvent *)event {
   NSLog(@"scrollWheel");
}

- (NSDragOperation)draggingEntered:(id <NSDraggingInfo>)sender {
    // HACK: We don't know what to say here because we don't know what the
    //       application wants to do with the paths
    return NSDragOperationGeneric;
}

- (BOOL)performDragOperation:(id <NSDraggingInfo>)sender {
 NSLog(@"performDragOp");
    return YES;
}

- (NSRect)firstRectForCharacterRange:(NSRange)range actualRange:(NSRangePointer)actualRange {
	NSLog(@"firstRectForCharacter");
return NSMakeRect(0, 0, 200, 200);
}

@end

void _platformCreateWindow() {
	id del = [[WindowDelegate alloc] init];
	NSRect contentRect = NSMakeRect(0, 0, 200, 200); //TODO: take in size
	id window = [[Window alloc] initWithContentRect:contentRect
	  styleMask:NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskTitled | NSWindowStyleMaskClosable | NSWindowStyleMaskResizable //TODO: get style mask
		backing:NSBackingStoreBuffered
		  defer:NO];
	id view = [[ContentView alloc] init];
	[window setContentView: view];
	[window makeFirstResponder: view];
	[window setTitle: @"im a title"];
	[window setDelegate: del];
	[window setRestorable: NO];
}

*/
import "C"
import (
	"github.com/pkg/errors"
	"log"
)

func PlatformInit() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.Wrap(p.(error), "Cocoa Init Failed:")
		}
	}()
	C._platformInit()
	return nil
}

func PlatformCreateWindow() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.Wrap(p.(error), "Cocoa Window Failed:")
		}
	}()
	C._platformCreateWindow()
	return nil
}

func main() {
	err := PlatformInit()
	if err != nil {
		log.Fatal(err)
	}
	PlatformCreateWindow()
}
