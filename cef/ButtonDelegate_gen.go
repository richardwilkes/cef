// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ButtonDelegate_gen.h"
	"C"
)

// ButtonDelegateProxy defines methods required for using ButtonDelegate.
type ButtonDelegateProxy interface {
	OnButtonPressed(self *ButtonDelegate, button *Button)
	OnButtonStateChanged(self *ButtonDelegate, button *Button)
}

// ButtonDelegate (cef_button_delegate_t from include/capi/views/cef_button_delegate_capi.h)
// Implement this structure to handle Button events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type ButtonDelegate C.cef_button_delegate_t

// NewButtonDelegate creates a new ButtonDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewButtonDelegate(proxy ButtonDelegateProxy) *ButtonDelegate {
	result := (*ButtonDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_button_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_button_delegate_proxy(result.toNative())
	}
	return result
}

func (d *ButtonDelegate) toNative() *C.cef_button_delegate_t {
	return (*C.cef_button_delegate_t)(d)
}

func lookupButtonDelegateProxy(obj *BaseRefCounted) ButtonDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ButtonDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ButtonDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ButtonDelegate) Base() *ViewDelegate {
	return (*ViewDelegate)(&d.base)
}

// OnButtonPressed (on_button_pressed)
// Called when |button| is pressed.
func (d *ButtonDelegate) OnButtonPressed(button *Button) {
	lookupButtonDelegateProxy(d.Base().Base()).OnButtonPressed(d, button)
}

//nolint:gocritic
//export gocef_button_delegate_on_button_pressed
func gocef_button_delegate_on_button_pressed(self *C.cef_button_delegate_t, button *C.cef_button_t) {
	me__ := (*ButtonDelegate)(self)
	proxy__ := lookupButtonDelegateProxy(me__.Base().Base())
	proxy__.OnButtonPressed(me__, (*Button)(button))
}

// OnButtonStateChanged (on_button_state_changed)
// Called when the state of |button| changes.
func (d *ButtonDelegate) OnButtonStateChanged(button *Button) {
	lookupButtonDelegateProxy(d.Base().Base()).OnButtonStateChanged(d, button)
}

//nolint:gocritic
//export gocef_button_delegate_on_button_state_changed
func gocef_button_delegate_on_button_state_changed(self *C.cef_button_delegate_t, button *C.cef_button_t) {
	me__ := (*ButtonDelegate)(self)
	proxy__ := lookupButtonDelegateProxy(me__.Base().Base())
	proxy__.OnButtonStateChanged(me__, (*Button)(button))
}
