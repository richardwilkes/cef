#import "app_darwin.h"
#import "include/capi/cef_app_capi.h"
#import "include/cef_application_mac.h"

@interface GoCEFApplication : NSApplication<CefAppProtocol>
{
	BOOL handlingSendEvent;
}
@end

@implementation GoCEFApplication

- (BOOL)isHandlingSendEvent {
	return handlingSendEvent;
}

- (void)setHandlingSendEvent:(BOOL)handling {
	handlingSendEvent = handling;
}

- (void)sendEvent:(NSEvent*)event {
	BOOL handling = handlingSendEvent;
	handlingSendEvent = YES;
	[super sendEvent:event];
	handlingSendEvent = handling;
}

@end

void gocef_instantiate_cef_application() {
	[GoCEFApplication sharedApplication];
}
