// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// This is a "blank identifier" assignment that acts as a compile-time check.
// The underscore (_) means "ignore this value" but the assignment still happens.
// This line ensures that our ComputeTaxFunction type actually implements
// the function.Function interface. If it doesn't, the Go compiler will
// give us an error, preventing bugs at compile time rather than runtime.
var _ function.Function = &ComputeTaxFunction{}

// This declares a new struct type called ComputeTaxFunction.
// A struct is like a template for creating objects that can hold data.
// In this case, our struct is empty (no fields), so it's just a placeholder
// that we'll attach methods to. Think of it like a blueprint for an object.
type ComputeTaxFunction struct{}

// This is a "constructor function" - a function that creates and returns
// a new instance of our ComputeTaxFunction struct.
// The function name starts with "New" which is a Go convention for constructors.
// It returns a function.Function interface type, which means any code that
// expects a function.Function can use the result of this function.
func NewComputeTaxFunction() function.Function {
	return &ComputeTaxFunction{} // The & creates a pointer to a new struct instance
}

func (f *ComputeTaxFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "compute_tax"
}

func (f *ComputeTaxFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Compute tax for coffee",
		Description: "Given a price and tax rate, return the total cost including tax.",
		Parameters: []function.Parameter{
			function.Float64Parameter{
				Name:        "price",
				Description: "Price of coffee item.",
			},
			function.Float64Parameter{
				Name:        "rate",
				Description: "Tax rate. 0.085 == 8.5%",
			},
		},
		Return: function.Float64Return{},
	}
}

func (f *ComputeTaxFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var price float64
	var rate float64
	var total float64

	// Read Terraform argument data into the variables
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &price, &rate))

	total = math.Round((price+price*rate)*100) / 100

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, total))
}
