// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#import "application_darwin.h"
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

void gocef_instantiate_application() {
	[GoCEFApplication sharedApplication];
}
