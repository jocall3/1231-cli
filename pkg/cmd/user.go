// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jocall3/1231-cli/internal/apiquery"
	"github.com/jocall3/1231-cli/internal/requestflag"
	"github.com/jocall3/go"
	"github.com/jocall3/go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usersLogin = cli.Command{
	Name:    "login",
	Usage:   "Authenticates a user and creates a secure session, returning access tokens. May\nrequire MFA depending on user settings.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "email",
			Usage:    "User's email address.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[any]{
			Name:     "password",
			Usage:    "User's password.",
			Required: true,
			BodyPath: "password",
		},
		&requestflag.Flag[any]{
			Name:     "mfa-code",
			Usage:    "Optional: Multi-factor authentication code, if required.",
			BodyPath: "mfaCode",
		},
	},
	Action:          handleUsersLogin,
	HideHelpCommand: true,
}

var usersRegister = requestflag.WithInnerFlags(cli.Command{
	Name:    "register",
	Usage:   "Registers a new user account with , initiating the onboarding process. Requires\nbasic user details.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "email",
			Usage:    "Email address for registration and login.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "Full name of the user.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "password",
			Usage:    "User's chosen password.",
			Required: true,
			BodyPath: "password",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "address",
			BodyPath: "address",
		},
		&requestflag.Flag[any]{
			Name:     "date-of-birth",
			Usage:    "Optional date of birth (YYYY-MM-DD).",
			BodyPath: "dateOfBirth",
		},
		&requestflag.Flag[any]{
			Name:     "phone",
			Usage:    "Optional phone number for MFA or recovery.",
			BodyPath: "phone",
		},
	},
	Action:          handleUsersRegister,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"address": {
		&requestflag.InnerFlag[any]{
			Name:       "address.city",
			InnerField: "city",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.country",
			InnerField: "country",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.state",
			InnerField: "state",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.street",
			InnerField: "street",
		},
		&requestflag.InnerFlag[any]{
			Name:       "address.zip",
			InnerField: "zip",
		},
	},
})

func handleUsersLogin(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserLoginParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Login(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users login", obj, format, transform)
}

func handleUsersRegister(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.UserRegisterParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Users.Register(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users register", obj, format, transform)
}
