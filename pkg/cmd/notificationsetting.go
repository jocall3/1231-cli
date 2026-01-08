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

var notificationsSettingsRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Retrieves the user's granular notification preferences across different channels\n(email, push, SMS, in-app) and event types.",
	Flags:           []cli.Flag{},
	Action:          handleNotificationsSettingsRetrieve,
	HideHelpCommand: true,
}

var notificationsSettingsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:  "update",
	Usage: "Updates the user's notification preferences, allowing control over channels,\nevent types, and quiet hours.",
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "channel-preferences",
			Usage:    "Updated preferences for notification delivery channels. Only provided fields are updated.",
			BodyPath: "channelPreferences",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "event-preferences",
			Usage:    "Updated preferences for different types of events. Only provided fields are updated.",
			BodyPath: "eventPreferences",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "quiet-hours",
			Usage:    "Updated settings for notification quiet hours. Only provided fields are updated.",
			BodyPath: "quietHours",
		},
	},
	Action:          handleNotificationsSettingsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"channel-preferences": {
		&requestflag.InnerFlag[any]{
			Name:       "channel-preferences.email",
			InnerField: "email",
		},
		&requestflag.InnerFlag[any]{
			Name:       "channel-preferences.in-app",
			InnerField: "inApp",
		},
		&requestflag.InnerFlag[any]{
			Name:       "channel-preferences.push",
			InnerField: "push",
		},
		&requestflag.InnerFlag[any]{
			Name:       "channel-preferences.sms",
			InnerField: "sms",
		},
	},
	"event-preferences": {
		&requestflag.InnerFlag[any]{
			Name:       "event-preferences.ai-insights",
			InnerField: "aiInsights",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event-preferences.budget-alerts",
			InnerField: "budgetAlerts",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event-preferences.promotional-offers",
			InnerField: "promotionalOffers",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event-preferences.security-alerts",
			InnerField: "securityAlerts",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event-preferences.transaction-alerts",
			InnerField: "transactionAlerts",
		},
	},
	"quiet-hours": {
		&requestflag.InnerFlag[any]{
			Name:       "quiet-hours.enabled",
			InnerField: "enabled",
		},
		&requestflag.InnerFlag[any]{
			Name:       "quiet-hours.end-time",
			InnerField: "endTime",
		},
		&requestflag.InnerFlag[any]{
			Name:       "quiet-hours.start-time",
			InnerField: "startTime",
		},
	},
})

func handleNotificationsSettingsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Notifications.Settings.Get(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications:settings retrieve", obj, format, transform)
}

func handleNotificationsSettingsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.NotificationSettingUpdateParams{}

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
	_, err = client.Notifications.Settings.Update(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications:settings update", obj, format, transform)
}
