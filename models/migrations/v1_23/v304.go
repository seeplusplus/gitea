// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_23 //nolint

import (
	"fmt"

	"xorm.io/xorm"
)

func AddShowPrivateActivityUserColumn(x *xorm.Engine) error {
	type User struct {
		ShowPrivateActivity bool `xorm:"NOT NULL DEFAULT false"`
	}

	if err := x.Sync(new(User)); err != nil {
		return fmt.Errorf("Sync: %w", err)
	}

	return nil
}
