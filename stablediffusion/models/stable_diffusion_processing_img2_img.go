// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StableDiffusionProcessingImg2Img StableDiffusionProcessingImg2Img
//
// swagger:model StableDiffusionProcessingImg2Img
type StableDiffusionProcessingImg2Img struct {

	// Alwayson Scripts
	AlwaysonScripts interface{} `json:"alwayson_scripts,omitempty"`

	// Batch Size
	BatchSize *int64 `json:"batch_size,omitempty"`

	// Cfg Scale
	CfgScale *float64 `json:"cfg_scale,omitempty"`

	// Comments
	Comments interface{} `json:"comments,omitempty"`

	// Denoising Strength
	DenoisingStrength *float64 `json:"denoising_strength,omitempty"`

	// Disable Extra Networks
	DisableExtraNetworks *bool `json:"disable_extra_networks,omitempty"`

	// Do Not Save Grid
	DoNotSaveGrid *bool `json:"do_not_save_grid,omitempty"`

	// Do Not Save Samples
	DoNotSaveSamples *bool `json:"do_not_save_samples,omitempty"`

	// Eta
	Eta float64 `json:"eta,omitempty"`

	// Height
	Height *int64 `json:"height,omitempty"`

	// Image Cfg Scale
	ImageCfgScale float64 `json:"image_cfg_scale,omitempty"`

	// Include Init Images
	IncludeInitImages *bool `json:"include_init_images,omitempty"`

	// Init Images
	InitImages []interface{} `json:"init_images"`

	// Initial Noise Multiplier
	InitialNoiseMultiplier float64 `json:"initial_noise_multiplier,omitempty"`

	// Inpaint Full Res
	InpaintFullRes *bool `json:"inpaint_full_res,omitempty"`

	// Inpaint Full Res Padding
	InpaintFullResPadding int64 `json:"inpaint_full_res_padding,omitempty"`

	// Inpainting Fill
	InpaintingFill int64 `json:"inpainting_fill,omitempty"`

	// Inpainting Mask Invert
	InpaintingMaskInvert int64 `json:"inpainting_mask_invert,omitempty"`

	// Latent Mask
	LatentMask string `json:"latent_mask,omitempty"`

	// Mask
	Mask string `json:"mask,omitempty"`

	// Mask Blur
	MaskBlur int64 `json:"mask_blur,omitempty"`

	// Mask Blur X
	MaskBlurx *int64 `json:"mask_blur_x,omitempty"`

	// Mask Blur Y
	MaskBlury *int64 `json:"mask_blur_y,omitempty"`

	// N Iter
	NIter *int64 `json:"n_iter,omitempty"`

	// Negative Prompt
	NegativePrompt string `json:"negative_prompt,omitempty"`

	// Override Settings
	OverrideSettings interface{} `json:"override_settings,omitempty"`

	// Override Settings Restore Afterwards
	OverrideSettingsRestoreAfterwards *bool `json:"override_settings_restore_afterwards,omitempty"`

	// Prompt
	Prompt string `json:"prompt,omitempty"`

	// Refiner Checkpoint
	RefinerCheckpoint string `json:"refiner_checkpoint,omitempty"`

	// Refiner Switch At
	RefinerSwitchAt float64 `json:"refiner_switch_at,omitempty"`

	// Resize Mode
	ResizeMode int64 `json:"resize_mode,omitempty"`

	// Restore Faces
	RestoreFaces bool `json:"restore_faces,omitempty"`

	// S Churn
	SChurn float64 `json:"s_churn,omitempty"`

	// S Min Uncond
	SMinUncond float64 `json:"s_min_uncond,omitempty"`

	// S Noise
	SNoise float64 `json:"s_noise,omitempty"`

	// S Tmax
	STmax float64 `json:"s_tmax,omitempty"`

	// S Tmin
	STmin float64 `json:"s_tmin,omitempty"`

	// Sampler Index
	SamplerIndex *string `json:"sampler_index,omitempty"`

	// Sampler Name
	SamplerName string `json:"sampler_name,omitempty"`

	// Scheduler Type
	Scheduler string `json:"scheduler,omitempty"`

	// Save Images
	SaveImages *bool `json:"save_images,omitempty"`

	// Script Args
	ScriptArgs []interface{} `json:"script_args"`

	// Script Name
	ScriptName string `json:"script_name,omitempty"`

	// Seed
	Seed *int64 `json:"seed,omitempty"`

	// Seed Resize From H
	SeedResizeFromh *int64 `json:"seed_resize_from_h,omitempty"`

	// Seed Resize From W
	SeedResizeFromw *int64 `json:"seed_resize_from_w,omitempty"`

	// Send Images
	SendImages *bool `json:"send_images,omitempty"`

	// Steps
	Steps *int64 `json:"steps,omitempty"`

	// Styles
	Styles []string `json:"styles"`

	// Subseed
	Subseed *int64 `json:"subseed,omitempty"`

	// Subseed Strength
	SubseedStrength float64 `json:"subseed_strength,omitempty"`

	// Tiling
	Tiling bool `json:"tiling,omitempty"`

	// Width
	Width *int64 `json:"width,omitempty"`
}

// Validate validates this stable diffusion processing img2 img
func (m *StableDiffusionProcessingImg2Img) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stable diffusion processing img2 img based on context it is used
func (m *StableDiffusionProcessingImg2Img) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StableDiffusionProcessingImg2Img) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StableDiffusionProcessingImg2Img) UnmarshalBinary(b []byte) error {
	var res StableDiffusionProcessingImg2Img
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
