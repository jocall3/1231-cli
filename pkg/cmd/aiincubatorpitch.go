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

var aiIncubatorPitchRetrieveDetails = cli.Command{
	Name:    "retrieve-details",
	Usage:   "Retrieves the granular AI-driven analysis, strategic feedback, market validation\nresults, and any outstanding questions from Quantum Weaver for a specific\nbusiness pitch.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "pitch-id",
			Required: true,
		},
	},
	Action:          handleAIIncubatorPitchRetrieveDetails,
	HideHelpCommand: true,
}

var aiIncubatorPitchSubmit = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit",
	Usage:   "Submits a detailed business plan to the Quantum Weaver AI for rigorous analysis,\nmarket validation, and seed funding consideration. This initiates the AI-driven\nincubation journey, aiming to transform innovative ideas into commercially\nsuccessful ventures.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "business-plan",
			Usage:    "The user's detailed narrative business plan (e.g., executive summary, vision, strategy).",
			Required: true,
			BodyPath: "businessPlan",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "financial-projections",
			Usage:    "Key financial metrics and projections for the next 3-5 years.",
			Required: true,
			BodyPath: "financialProjections",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "founding-team",
			Usage:    "Key profiles and expertise of the founding team members.",
			Required: true,
			BodyPath: "foundingTeam",
		},
		&requestflag.Flag[any]{
			Name:     "market-opportunity",
			Usage:    "Detailed analysis of the target market, problem statement, and proposed solution's unique value proposition.",
			Required: true,
			BodyPath: "marketOpportunity",
		},
	},
	Action:          handleAIIncubatorPitchSubmit,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"financial-projections": {
		&requestflag.InnerFlag[any]{
			Name:       "financial-projections.profitability-estimate",
			Usage:      "Estimated time to profitability.",
			InnerField: "profitabilityEstimate",
		},
		&requestflag.InnerFlag[any]{
			Name:       "financial-projections.projection-years",
			Usage:      "Number of years for financial projections.",
			InnerField: "projectionYears",
		},
		&requestflag.InnerFlag[[]any]{
			Name:       "financial-projections.revenue-forecast",
			InnerField: "revenueForecast",
		},
		&requestflag.InnerFlag[any]{
			Name:       "financial-projections.seed-round-amount",
			Usage:      "Requested seed funding in USD.",
			InnerField: "seedRoundAmount",
		},
		&requestflag.InnerFlag[any]{
			Name:       "financial-projections.valuation-pre-money",
			Usage:      "Pre-money valuation in USD.",
			InnerField: "valuationPreMoney",
		},
	},
	"founding-team": {
		&requestflag.InnerFlag[any]{
			Name:       "founding-team.experience",
			Usage:      "Relevant experience.",
			InnerField: "experience",
		},
		&requestflag.InnerFlag[any]{
			Name:       "founding-team.name",
			Usage:      "Name of the team member.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[any]{
			Name:       "founding-team.role",
			Usage:      "Role of the team member.",
			InnerField: "role",
		},
	},
})

var aiIncubatorPitchSubmitFeedback = requestflag.WithInnerFlags(cli.Command{
	Name:    "submit-feedback",
	Usage:   "Allows the entrepreneur to respond to specific questions or provide additional\ndetails requested by Quantum Weaver, moving the pitch forward in the incubation\nprocess.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "pitch-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "answer",
			BodyPath: "answers",
		},
		&requestflag.Flag[any]{
			Name:     "feedback",
			Usage:    "General textual feedback or additional details for Quantum Weaver.",
			BodyPath: "feedback",
		},
	},
	Action:          handleAIIncubatorPitchSubmitFeedback,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"answer": {
		&requestflag.InnerFlag[any]{
			Name:       "answer.answer",
			Usage:      "The answer to the specific question.",
			InnerField: "answer",
		},
		&requestflag.InnerFlag[any]{
			Name:       "answer.question-id",
			Usage:      "The ID of the question being answered.",
			InnerField: "questionId",
		},
	},
})

func handleAIIncubatorPitchRetrieveDetails(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pitch-id") && len(unusedArgs) > 0 {
		cmd.Set("pitch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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
	_, err = client.AI.Incubator.Pitch.GetDetails(ctx, interface{}(cmd.Value("pitch-id").(any)), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:incubator:pitch retrieve-details", obj, format, transform)
}

func handleAIIncubatorPitchSubmit(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIIncubatorPitchSubmitParams{}

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
	_, err = client.AI.Incubator.Pitch.Submit(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:incubator:pitch submit", obj, format, transform)
}

func handleAIIncubatorPitchSubmitFeedback(ctx context.Context, cmd *cli.Command) error {
	client := jocall3.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pitch-id") && len(unusedArgs) > 0 {
		cmd.Set("pitch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := jocall3.AIIncubatorPitchSubmitFeedbackParams{}

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
	_, err = client.AI.Incubator.Pitch.SubmitFeedback(
		ctx,
		interface{}(cmd.Value("pitch-id").(any)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "ai:incubator:pitch submit-feedback", obj, format, transform)
}
