// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package webauthnx

import (
	"crypto/sha512"
	_ "embed"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/ory/x/stringsx"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/text"
	"github.com/ory/kratos/ui/node"
)

func NewWebAuthnConnectionTrigger(options string) *node.Node {
	return node.NewInputField(node.WebAuthnRegisterTrigger, "", node.WebAuthnGroup,
		node.InputAttributeTypeButton, node.WithInputAttributes(func(a *node.InputAttributes) {
			a.OnClick = "window.__oryWebAuthnRegistration(" + options + ")"
		}))
}

func NewWebAuthnScript(base *url.URL) *node.Node {
	src := urlx.AppendPaths(base, ScriptURL).String()
	integrity := sha512.Sum512(jsOnLoad)
	return node.NewScriptField(
		node.WebAuthnScript,
		src,
		node.WebAuthnGroup,
		fmt.Sprintf("sha512-%s", base64.StdEncoding.EncodeToString(integrity[:])),
	)
}

func NewCreatePasskeyScript(base *url.URL) *node.Node {
	src := urlx.AppendPaths(base, CreatePasskeyScriptURL).String()
	integrity := sha512.Sum512(createPasskeyJS)
	return node.NewScriptField(
		node.CreatePasskey,
		src,
		node.WebAuthnGroup,
		fmt.Sprintf("sha512-%s", base64.StdEncoding.EncodeToString(integrity[:])),
	)
}

func NewWebAuthnConnectionInput() *node.Node {
	return node.NewInputField(node.WebAuthnRegister, "", node.WebAuthnGroup,
		node.InputAttributeTypeHidden)
}

func NewWebAuthnLoginTrigger(options string) *node.Node {
	return node.NewInputField(node.WebAuthnLoginTrigger, "", node.WebAuthnGroup,
		node.InputAttributeTypeButton, node.WithInputAttributes(func(a *node.InputAttributes) {
			a.OnClick = "window.__oryWebAuthnLogin(" + options + ")"
		}))
}

func NewWebAuthnLoginInput() *node.Node {
	return node.NewInputField(node.WebAuthnLogin, "", node.WebAuthnGroup,
		node.InputAttributeTypeHidden)
}

func NewWebAuthnConnectionName() *node.Node {
	return node.NewInputField(node.WebAuthnRegisterDisplayName, "", node.WebAuthnGroup, node.InputAttributeTypeText).
		WithMetaLabel(text.NewInfoSelfServiceRegisterWebAuthnDisplayName())
}

func NewWebAuthnUnlink(c *identity.CredentialWebAuthn) *node.Node {
	return node.NewInputField(node.WebAuthnRemove, fmt.Sprintf("%x", c.ID), node.WebAuthnGroup,
		node.InputAttributeTypeSubmit).
		WithMetaLabel(text.NewInfoSelfServiceRemoveWebAuthn(stringsx.Coalesce(c.DisplayName, "unnamed"), c.AddedAt))
}