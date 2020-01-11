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
	// #include "TextfieldDelegate_gen.h"
	"C"
)

// TextfieldDelegateProxy defines methods required for using TextfieldDelegate.
type TextfieldDelegateProxy interface {
	OnKeyEvent(self *TextfieldDelegate, textfield *Textfield, event *KeyEvent) int32
	OnAfterUserAction(self *TextfieldDelegate, textfield *Textfield)
}

// TextfieldDelegate (cef_textfield_delegate_t from include/capi/views/cef_textfield_delegate_capi.h)
// Implement this structure to handle Textfield events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type TextfieldDelegate C.cef_textfield_delegate_t

// NewTextfieldDelegate creates a new TextfieldDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewTextfieldDelegate(proxy TextfieldDelegateProxy) *TextfieldDelegate {
	result := (*TextfieldDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_textfield_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_textfield_delegate_proxy(result.toNative())
	}
	return result
}

func (d *TextfieldDelegate) toNative() *C.cef_textfield_delegate_t {
	return (*C.cef_textfield_delegate_t)(d)
}

func lookupTextfieldDelegateProxy(obj *BaseRefCounted) TextfieldDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(TextfieldDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type TextfieldDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *TextfieldDelegate) Base() *ViewDelegate {
	return (*ViewDelegate)(&d.base)
}

// OnKeyEvent (on_key_event)
// Called when |textfield| recieves a keyboard event. |event| contains
// information about the keyboard event. Return true (1) if the keyboard event
// was handled or false (0) otherwise for default handling.
func (d *TextfieldDelegate) OnKeyEvent(textfield *Textfield, event *KeyEvent) int32 {
	return lookupTextfieldDelegateProxy(d.Base().Base()).OnKeyEvent(d, textfield, event)
}

//nolint:gocritic
//export gocef_textfield_delegate_on_key_event
func gocef_textfield_delegate_on_key_event(self *C.cef_textfield_delegate_t, textfield *C.cef_textfield_t, event *C.cef_key_event_t) C.int {
	me__ := (*TextfieldDelegate)(self)
	proxy__ := lookupTextfieldDelegateProxy(me__.Base().Base())
	event_ := event.toGo()
	return C.int(proxy__.OnKeyEvent(me__, (*Textfield)(textfield), event_))
}

// OnAfterUserAction (on_after_user_action)
// Called after performing a user action that may change |textfield|.
func (d *TextfieldDelegate) OnAfterUserAction(textfield *Textfield) {
	lookupTextfieldDelegateProxy(d.Base().Base()).OnAfterUserAction(d, textfield)
}

//nolint:gocritic
//export gocef_textfield_delegate_on_after_user_action
func gocef_textfield_delegate_on_after_user_action(self *C.cef_textfield_delegate_t, textfield *C.cef_textfield_t) {
	me__ := (*TextfieldDelegate)(self)
	proxy__ := lookupTextfieldDelegateProxy(me__.Base().Base())
	proxy__.OnAfterUserAction(me__, (*Textfield)(textfield))
}
