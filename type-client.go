package ga

//WARNING: This file was generated. Do not edit.

import "net/url"
import "errors"

func (c *Client) setType(h hitType) {
	switch h.(type) {
	case *Event:
		c.hitType = "event"
	case *Exception:
		c.hitType = "exception"
	case *Item:
		c.hitType = "item"
	case *Pageview:
		c.hitType = "pageview"
	case *Screenview:
		c.hitType = "screenview"
	case *Social:
		c.hitType = "social"
	case *Timing:
		c.hitType = "timing"
	case *Transaction:
		c.hitType = "transaction"
	}
}

//Client Hit Type
type Client struct {
	//Use TLS when Send()ing
	UseTLS bool
	// The Protocol version. The current value is '1'. This will
	// only change when there are changes made that are not backwards
	// compatible.
	ProtocolVersion string
	// The tracking ID / web property ID. The format is UA-XXXX-Y.
	// All collected data is associated by this ID.
	TrackingID string
	// When present, the IP address of the sender will be anonymized.
	// For example, the IP will be anonymized if any of the following
	// parameters are present in the payload: &aip=, &aip=0, or
	// &aip=1
	AnonymizeIP string
	// Used to collect offline / latent hits. The value represents
	// the time delta (in milliseconds) between when the hit being
	// reported occurred and the time the hit was sent. The value
	// must be greater than or equal to 0. Values greater than
	// four hours may lead to hits not being processed.
	QueueTime string
	// Used to send a random number in GET requests to ensure browsers
	// and proxies don't cache hits. It should be sent as the final
	// parameter of the request since we've seen some 3rd party
	// internet filtering software add additional parameters to
	// HTTP requests incorrectly. This value is not used in reporting.
	CacheBuster string
	// This anonymously identifies a particular user, device, or
	// browser instance. For the web, this is generally stored
	// as a first-party cookie with a two-year expiration. For
	// mobile apps, this is randomly generated for each particular
	// instance of an application install. The value of this field
	// should be a random UUID (version 4) as described in http://www.ietf.org/rfc/rfc4122.txt
	ClientID string
	// This is intended to be a known identifier for a user provided
	// by the site owner/tracking library user. It may not itself
	// be PII (personally identifiable information). The value
	// should never be persisted in GA cookies or other Analytics
	// provided storage.
	UserID string
	// Used to control the session duration. A value of 'start'
	// forces a new session to start with this hit and 'end' forces
	// the current session to end with this hit. All other values
	// are ignored.
	SessionControl string
	// The IP address of the user. This should be a valid IP address
	// in IPv4 or IPv6 format. It will always be anonymized just
	// as though &aip (anonymize IP) had been used.
	IPOverride string
	// The User Agent of the browser. Note that Google has libraries
	// to identify real user agents. Hand crafting your own agent
	// could break at any time.
	UserAgentOverride string
	// The geographical location of the user. The geographical
	// ID should be a two letter country code or a criteria ID
	// representing a city or region (see http://developers.google.com/analytics/devguides/collection/protocol/v1/geoid).
	// This parameter takes precedent over any location derived
	// from IP address, including the IP Override parameter. An
	// invalid code will result in geographical dimensions to be
	// set to '(not set)'.
	GeographicalOverride string
	// Specifies which referral source brought traffic to a website.
	// This value is also used to compute the traffic source. The
	// format of this value is a URL.
	DocumentReferrer string
	// Specifies the campaign name.
	CampaignName string
	// Specifies the campaign source.
	CampaignSource string
	// Specifies the campaign medium.
	CampaignMedium string
	// Specifies the campaign keyword.
	CampaignKeyword string
	// Specifies the campaign content.
	CampaignContent string
	// Specifies the campaign ID.
	CampaignID string
	// Specifies the Google AdWords Id.
	GoogleAdWordsID string
	// Specifies the Google Display Ads Id.
	GoogleDisplayAdsID string
	// Specifies the screen resolution.
	ScreenResolution string
	// Specifies the viewable area of the browser / device.
	ViewportSize string
	// Specifies the character set used to encode the page / document.
	DocumentEncoding string
	// Specifies the screen color depth.
	ScreenColors string
	// Specifies the language.
	UserLanguage string
	// Specifies whether Java was enabled.
	JavaEnabled string
	// Specifies the flash version.
	FlashVersion string
	// The type of hit. Must be one of 'pageview', 'screenview',
	// 'event', 'transaction', 'item', 'social', 'exception', 'timing'.
	hitType string
	// Specifies that a hit be considered non-interactive.
	NonInteractionHit string
	// Use this parameter to send the full URL (document location)
	// of the page on which content resides. You can use the &dh
	// and &dp parameters to override the hostname and path + query
	// portions of the document location, accordingly. The JavaScript
	// clients determine this parameter using the concatenation
	// of the document.location.origin + document.location.pathname
	// + document.location.search browser parameters. Be sure to
	// remove any user authentication or other private information
	// from the URL if present. For 'pageview' hits, either &dl
	// or both &dh and &dp have to be specified for the hit to
	// be valid.
	DocumentLocationURL string
	// Specifies the hostname from which content was hosted.
	DocumentHostName string
	// The path portion of the page URL. Should begin with '/'.
	// For 'pageview' hits, either &dl or both &dh and &dp have
	// to be specified for the hit to be valid.
	DocumentPath string
	// The title of the page / document.
	DocumentTitle string
	// If not specified, this will default to the unique URL of
	// the page by either using the &dl parameter as-is or assembling
	// it from &dh and &dp. App tracking makes use of this for
	// the 'Screen Name' of the screenview hit.
	ScreenName string
	// The ID of a clicked DOM element, used to disambiguate multiple
	// links to the same URL in In-Page Analytics reports when
	// Enhanced Link Attribution is enabled for the property.
	LinkID string
	// Specifies the application name.
	ApplicationName string
	// Application identifier.
	ApplicationID string
	// Specifies the application version.
	ApplicationVersion string
	// Application installer identifier.
	ApplicationInstallerID string
	// The SKU of the product. Product index must be a positive
	// integer between 1 and 200, inclusive. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	ProductSKU string
	// The name of the product. Product index must be a positive
	// integer between 1 and 200, inclusive. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	ProductName string
	// The brand associated with the product. Product index must
	// be a positive integer between 1 and 200, inclusive. For
	// analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductBrand string
	// The category to which the product belongs. Product index
	// must be a positive integer between 1 and 200, inclusive.
	// The product category parameter can be hierarchical. Use
	// / as a delimiter to specify up to 5-levels of hierarchy.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductCategory string
	// The variant of the product. Product index must be a positive
	// integer between 1 and 200, inclusive. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	ProductVariant string
	// The price of a product. Product index must be a positive
	// integer between 1 and 200, inclusive. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	ProductPrice string
	// The quantity of a product. Product index must be a positive
	// integer between 1 and 200, inclusive. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	ProductQuantity string
	// The coupon code associated with a product. Product index
	// must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductCouponCode string
	// The product's position in a list or collection. Product
	// index must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductPosition string
	// A product-level custom dimension where dimension index is
	// a positive integer between 1 and 200, inclusive. Product
	// index must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductCustomDimension string
	// A product-level custom metric where metric index is a positive
	// integer between 1 and 200, inclusive. Product index must
	// be a positive integer between 1 and 200, inclusive. For
	// analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductCustomMetric string
	// The role of the products included in a hit. If a product
	// action is not specified, all product definitions included
	// with the hit will be ignored. Must be one of: detail, click,
	// add, remove, checkout, checkout_option, purchase, refund.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductAction string
	// The transaction ID. This is an additional parameter that
	// can be sent when Product Action is set to 'purchase' or
	// 'refund'. For analytics.js the Enhanced Ecommerce plugin
	// must be installed before using this field.
	TransactionID string
	// The store or affiliation from which this transaction occurred.
	// This is an additional parameter that can be sent when Product
	// Action is set to 'purchase' or 'refund'. For analytics.js
	// the Enhanced Ecommerce plugin must be installed before using
	// this field.
	Affiliation string
	// The total value of the transaction, including tax and shipping.
	// If not sent, this value will be automatically calculated
	// using the product quantity and price fields of all products
	// in the same hit. This is an additional parameter that can
	// be sent when Product Action is set to 'purchase' or 'refund'.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	Revenue string
	// The total tax associated with the transaction. This is an
	// additional parameter that can be sent when Product Action
	// is set to 'purchase' or 'refund'. For analytics.js the Enhanced
	// Ecommerce plugin must be installed before using this field.
	Tax string
	// The shipping cost associated with the transaction. This
	// is an additional parameter that can be sent when Product
	// Action is set to 'purchase' or 'refund'. For analytics.js
	// the Enhanced Ecommerce plugin must be installed before using
	// this field.
	Shipping string
	// The transaction coupon redeemed with the transaction. This
	// is an additional parameter that can be sent when Product
	// Action is set to 'purchase' or 'refund'. For analytics.js
	// the Enhanced Ecommerce plugin must be installed before using
	// this field.
	CouponCode string
	// The list or collection from which a product action occurred.
	// This is an additional parameter that can be sent when Product
	// Action is set to 'detail' or 'click'. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	ProductActionList string
	// The step number in a checkout funnel. This is an additional
	// parameter that can be sent when Product Action is set to
	// 'checkout'. For analytics.js the Enhanced Ecommerce plugin
	// must be installed before using this field.
	CheckoutStep string
	// Additional information about a checkout step. This is an
	// additional parameter that can be sent when Product Action
	// is set to 'checkout'. For analytics.js the Enhanced Ecommerce
	// plugin must be installed before using this field.
	CheckoutStepOption string
	// The list or collection to which a product belongs. Impression
	// List index must be a positive integer between 1 and 200,
	// inclusive. For analytics.js the Enhanced Ecommerce plugin
	// must be installed before using this field.
	ProductImpressionListName string
	// The product ID or SKU. Impression List index must be a positive
	// integer between 1 and 200, inclusive. Product index must
	// be a positive integer between 1 and 200, inclusive. For
	// analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductImpressionSKU string
	// The name of the product. Impression List index must be a
	// positive integer between 1 and 200, inclusive. Product index
	// must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductImpressionName string
	// The brand associated with the product. Impression List index
	// must be a positive integer between 1 and 200, inclusive.
	// Product index must be a positive integer between 1 and 200,
	// inclusive. For analytics.js the Enhanced Ecommerce plugin
	// must be installed before using this field.
	ProductImpressionBrand string
	// The category to which the product belongs. Impression List
	// index must be a positive integer between 1 and 200, inclusive.
	// Product index must be a positive integer between 1 and 200,
	// inclusive. For analytics.js the Enhanced Ecommerce plugin
	// must be installed before using this field.
	ProductImpressionCategory string
	// The variant of the product. Impression List index must be
	// a positive integer between 1 and 200, inclusive. Product
	// index must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductImpressionVariant string
	// The product's position in a list or collection. Impression
	// List index must be a positive integer between 1 and 200,
	// inclusive. Product index must be a positive integer between
	// 1 and 200, inclusive. For analytics.js the Enhanced Ecommerce
	// plugin must be installed before using this field.
	ProductImpressionPosition string
	// The price of a product. Impression List index must be a
	// positive integer between 1 and 200, inclusive. Product index
	// must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	ProductImpressionPrice string
	// A product-level custom dimension where dimension index is
	// a positive integer between 1 and 200, inclusive. Impression
	// List index must be a positive integer between 1 and 200,
	// inclusive. Product index must be a positive integer between
	// 1 and 200, inclusive. For analytics.js the Enhanced Ecommerce
	// plugin must be installed before using this field.
	ProductImpressionCustomDimension string
	// A product-level custom metric where metric index is a positive
	// integer between 1 and 200, inclusive. Impression List index
	// must be a positive integer between 1 and 200, inclusive.
	// Product index must be a positive integer between 1 and 200,
	// inclusive. For analytics.js the Enhanced Ecommerce plugin
	// must be installed before using this field.
	ProductImpressionCustomMetric string
	// The promotion ID. Promotion index must be a positive integer
	// between 1 and 200, inclusive. For analytics.js the Enhanced
	// Ecommerce plugin must be installed before using this field.
	PromotionID string
	// The name of the promotion. Promotion index must be a positive
	// integer between 1 and 200, inclusive. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	PromotionName string
	// The creative associated with the promotion. Promotion index
	// must be a positive integer between 1 and 200, inclusive.
	// For analytics.js the Enhanced Ecommerce plugin must be installed
	// before using this field.
	PromotionCreative string
	// The position of the creative. Promotion index must be a
	// positive integer between 1 and 200, inclusive. For analytics.js
	// the Enhanced Ecommerce plugin must be installed before using
	// this field.
	PromotionPosition string
	// Specifies the role of the promotions included in a hit.
	// If a promotion action is not specified, the default promotion
	// action, 'view', is assumed. To measure a user click on a
	// promotion set this to 'promo_click'. For analytics.js the
	// Enhanced Ecommerce plugin must be installed before using
	// this field.
	PromotionAction string
	// Each custom dimension has an associated index. There is
	// a maximum of 20 custom dimensions (200 for Premium accounts).
	// The dimension index must be a positive integer between 1
	// and 200, inclusive.
	CustomDimension string
	// Each custom metric has an associated index. There is a maximum
	// of 20 custom metrics (200 for Premium accounts). The metric
	// index must be a positive integer between 1 and 200, inclusive.
	CustomMetric string
	// This parameter specifies that this user has been exposed
	// to an experiment with the given ID. It should be sent in
	// conjunction with the Experiment Variant parameter.
	ExperimentID string
	// This parameter specifies that this user has been exposed
	// to a particular variation of an experiment. It should be
	// sent in conjunction with the Experiment ID parameter.
	ExperimentVariant string
	// DimensionIndex is required by other properties
	DimensionIndex string
	// ListIndex is required by other properties
	ListIndex string
	// MetricIndex is required by other properties
	MetricIndex string
	// ProductIndex is required by other properties
	ProductIndex string
	// PromoIndex is required by other properties
	PromoIndex string
}

func (h *Client) addFields(v url.Values) error {
	if h.ProtocolVersion != "" {
		v.Add("v", h.ProtocolVersion)
	} else {
		return errors.New("Required field 'ProtocolVersion' is missing")
	}
	if h.TrackingID != "" {
		v.Add("tid", h.TrackingID)
	} else {
		return errors.New("Required field 'TrackingID' is missing")
	}
	if h.AnonymizeIP != "" {
		v.Add("aip", h.AnonymizeIP)
	}
	if h.QueueTime != "" {
		v.Add("qt", h.QueueTime)
	}
	if h.CacheBuster != "" {
		v.Add("z", h.CacheBuster)
	}
	if h.ClientID != "" {
		v.Add("cid", h.ClientID)
	} else {
		return errors.New("Required field 'ClientID' is missing")
	}
	if h.UserID != "" {
		v.Add("uid", h.UserID)
	}
	if h.SessionControl != "" {
		v.Add("sc", h.SessionControl)
	}
	if h.IPOverride != "" {
		v.Add("uip", h.IPOverride)
	}
	if h.UserAgentOverride != "" {
		v.Add("ua", h.UserAgentOverride)
	}
	if h.GeographicalOverride != "" {
		v.Add("geoid", h.GeographicalOverride)
	}
	if h.DocumentReferrer != "" {
		v.Add("dr", h.DocumentReferrer)
	}
	if h.CampaignName != "" {
		v.Add("cn", h.CampaignName)
	}
	if h.CampaignSource != "" {
		v.Add("cs", h.CampaignSource)
	}
	if h.CampaignMedium != "" {
		v.Add("cm", h.CampaignMedium)
	}
	if h.CampaignKeyword != "" {
		v.Add("ck", h.CampaignKeyword)
	}
	if h.CampaignContent != "" {
		v.Add("cc", h.CampaignContent)
	}
	if h.CampaignID != "" {
		v.Add("ci", h.CampaignID)
	}
	if h.GoogleAdWordsID != "" {
		v.Add("gclid", h.GoogleAdWordsID)
	}
	if h.GoogleDisplayAdsID != "" {
		v.Add("dclid", h.GoogleDisplayAdsID)
	}
	if h.ScreenResolution != "" {
		v.Add("sr", h.ScreenResolution)
	}
	if h.ViewportSize != "" {
		v.Add("vp", h.ViewportSize)
	}
	if h.DocumentEncoding != "" {
		v.Add("de", h.DocumentEncoding)
	}
	if h.ScreenColors != "" {
		v.Add("sd", h.ScreenColors)
	}
	if h.UserLanguage != "" {
		v.Add("ul", h.UserLanguage)
	}
	if h.JavaEnabled != "" {
		v.Add("je", h.JavaEnabled)
	}
	if h.FlashVersion != "" {
		v.Add("fl", h.FlashVersion)
	}
	if h.hitType != "" {
		v.Add("t", h.hitType)
	} else {
		return errors.New("Required field 'hitType' is missing")
	}
	if h.NonInteractionHit != "" {
		v.Add("ni", h.NonInteractionHit)
	}
	if h.DocumentLocationURL != "" {
		v.Add("dl", h.DocumentLocationURL)
	}
	if h.DocumentHostName != "" {
		v.Add("dh", h.DocumentHostName)
	}
	if h.DocumentPath != "" {
		v.Add("dp", h.DocumentPath)
	}
	if h.DocumentTitle != "" {
		v.Add("dt", h.DocumentTitle)
	}
	if h.ScreenName != "" {
		v.Add("cd", h.ScreenName)
	}
	if h.LinkID != "" {
		v.Add("linkid", h.LinkID)
	}
	if h.ApplicationName != "" {
		v.Add("an", h.ApplicationName)
	}
	if h.ApplicationID != "" {
		v.Add("aid", h.ApplicationID)
	}
	if h.ApplicationVersion != "" {
		v.Add("av", h.ApplicationVersion)
	}
	if h.ApplicationInstallerID != "" {
		v.Add("aiid", h.ApplicationInstallerID)
	}
	if h.ProductSKU != "" {
		v.Add("pr" + h.ProductIndex + "id", h.ProductSKU)
	}
	if h.ProductName != "" {
		v.Add("pr" + h.ProductIndex + "nm", h.ProductName)
	}
	if h.ProductBrand != "" {
		v.Add("pr" + h.ProductIndex + "br", h.ProductBrand)
	}
	if h.ProductCategory != "" {
		v.Add("pr" + h.ProductIndex + "ca", h.ProductCategory)
	}
	if h.ProductVariant != "" {
		v.Add("pr" + h.ProductIndex + "va", h.ProductVariant)
	}
	if h.ProductPrice != "" {
		v.Add("pr" + h.ProductIndex + "pr", h.ProductPrice)
	}
	if h.ProductQuantity != "" {
		v.Add("pr" + h.ProductIndex + "qt", h.ProductQuantity)
	}
	if h.ProductCouponCode != "" {
		v.Add("pr" + h.ProductIndex + "cc", h.ProductCouponCode)
	}
	if h.ProductPosition != "" {
		v.Add("pr" + h.ProductIndex + "ps", h.ProductPosition)
	}
	if h.ProductCustomDimension != "" {
		v.Add("pr" + h.ProductIndex + "cd" + h.DimensionIndex + "", h.ProductCustomDimension)
	}
	if h.ProductCustomMetric != "" {
		v.Add("pr" + h.ProductIndex + "cm" + h.MetricIndex + "", h.ProductCustomMetric)
	}
	if h.ProductAction != "" {
		v.Add("pa", h.ProductAction)
	}
	if h.TransactionID != "" {
		v.Add("ti", h.TransactionID)
	}
	if h.Affiliation != "" {
		v.Add("ta", h.Affiliation)
	}
	if h.Revenue != "" {
		v.Add("tr", h.Revenue)
	}
	if h.Tax != "" {
		v.Add("tt", h.Tax)
	}
	if h.Shipping != "" {
		v.Add("ts", h.Shipping)
	}
	if h.CouponCode != "" {
		v.Add("tcc", h.CouponCode)
	}
	if h.ProductActionList != "" {
		v.Add("pal", h.ProductActionList)
	}
	if h.CheckoutStep != "" {
		v.Add("cos", h.CheckoutStep)
	}
	if h.CheckoutStepOption != "" {
		v.Add("col", h.CheckoutStepOption)
	}
	if h.ProductImpressionListName != "" {
		v.Add("il" + h.ListIndex + "nm", h.ProductImpressionListName)
	}
	if h.ProductImpressionSKU != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "id", h.ProductImpressionSKU)
	}
	if h.ProductImpressionName != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "nm", h.ProductImpressionName)
	}
	if h.ProductImpressionBrand != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "br", h.ProductImpressionBrand)
	}
	if h.ProductImpressionCategory != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "ca", h.ProductImpressionCategory)
	}
	if h.ProductImpressionVariant != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "va", h.ProductImpressionVariant)
	}
	if h.ProductImpressionPosition != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "ps", h.ProductImpressionPosition)
	}
	if h.ProductImpressionPrice != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "pr", h.ProductImpressionPrice)
	}
	if h.ProductImpressionCustomDimension != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "cd" + h.DimensionIndex + "", h.ProductImpressionCustomDimension)
	}
	if h.ProductImpressionCustomMetric != "" {
		v.Add("il" + h.ListIndex + "pi" + h.ProductIndex + "cm" + h.MetricIndex + "", h.ProductImpressionCustomMetric)
	}
	if h.PromotionID != "" {
		v.Add("promo" + h.PromoIndex + "id", h.PromotionID)
	}
	if h.PromotionName != "" {
		v.Add("promo" + h.PromoIndex + "nm", h.PromotionName)
	}
	if h.PromotionCreative != "" {
		v.Add("promo" + h.PromoIndex + "cr", h.PromotionCreative)
	}
	if h.PromotionPosition != "" {
		v.Add("promo" + h.PromoIndex + "ps", h.PromotionPosition)
	}
	if h.PromotionAction != "" {
		v.Add("promoa", h.PromotionAction)
	}
	if h.CustomDimension != "" {
		v.Add("cd" + h.DimensionIndex + "", h.CustomDimension)
	}
	if h.CustomMetric != "" {
		v.Add("cm" + h.MetricIndex + "", h.CustomMetric)
	}
	if h.ExperimentID != "" {
		v.Add("xid", h.ExperimentID)
	}
	if h.ExperimentVariant != "" {
		v.Add("xvar", h.ExperimentVariant)
	}
	return nil
}