package cmp

import (
	"fmt"
	"net/url"
	"strings"
)

type Companies map[string][]Company

type Company struct {
	Name          string   `json:"name"`
	URL           string   `json:"url"`
	CareerPageURL string   `json:"career_page_url"`
	Type          string   `json:"type,omitempty"`
	Categories    []string `json:"categories,omitempty"`
	RemotePolicy  string   `json:"remote_policy,omitempty"`
	HiringPolicy  string   `json:"hiring_policy,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}

var (
	allowedCategories = map[string]bool{
		"cloud_software":    true,
		"design_ux":         true,
		"marketing_writing": true,
		"hr":                true,
	}
	allowedHiringPolicies = map[string]bool{
		"-":            true,
		"Contract":     true,
		"Direct":       true,
		"Intermediary": true,
	}
	allowedRemotePolicies = map[string]bool{
		"-":        true,
		"Full":     true,
		"Hybrid":   true,
		"Optional": true,
	}
	allowedCompanyTypes = map[string]bool{
		"B2B":        true,
		"B2C":        true,
		"Consulting": true,
		"Product":    true,
	}
)

func (c Company) GetTagsString() string {
	return strings.Join(c.Tags, " - ")
}

func (c Company) Validate() error {
	for _, category := range c.Categories {
		if !allowedCategories[category] {
			return fmt.Errorf("%s | Category %s not allowed", c.Name, category)
		}
	}

	if !allowedHiringPolicies[c.HiringPolicy] {
		return fmt.Errorf("%s | Hiring Policy %s not allowed", c.Name, c.HiringPolicy)
	}

	if !allowedRemotePolicies[c.RemotePolicy] {
		return fmt.Errorf("%s | Remote Policy %s not allowed", c.Name, c.RemotePolicy)
	}

	if !allowedCompanyTypes[c.Type] {
		if len(c.Categories) > 0 && c.Categories[0] == "cloud_software" {
			return fmt.Errorf("%s | Type %s not allowed", c.Name, c.Type)
		}
	}

	if len(c.Tags) > 20 {
		return fmt.Errorf("Company tags must be less than 20")
	}

	_, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return fmt.Errorf("Company URL is not valid. ", err)
	}

	_, err = url.ParseRequestURI(c.CareerPageURL)
	if err != nil {
		return fmt.Errorf("Company Career Page URL is not valid. ", err)
	}

	return nil
}
