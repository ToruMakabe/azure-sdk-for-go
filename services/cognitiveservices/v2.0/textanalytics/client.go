// Package textanalytics implements the Azure ARM Textanalytics service API version v2.0.
//
// The Text Analytics API is a suite of text analytics web services built with best-in-class Microsoft machine learning
// algorithms. The API can be used to analyze unstructured text for tasks such as sentiment analysis, key phrase
// extraction and language detection. No training data is needed to use this API; just bring your text data. This API
// uses advanced natural language processing techniques to deliver best in class predictions. Further documentation can
// be found in https://docs.microsoft.com/en-us/azure/cognitive-services/text-analytics/overview
package textanalytics

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// BaseClient is the base client for Textanalytics.
type BaseClient struct {
	autorest.Client
	AzureRegion AzureRegions
}

// New creates an instance of the BaseClient client.
func New(azureRegion AzureRegions) BaseClient {
	return NewWithoutDefaults(azureRegion)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(azureRegion AzureRegions) BaseClient {
	return BaseClient{
		Client:      autorest.NewClientWithUserAgent(UserAgent()),
		AzureRegion: azureRegion,
	}
}

// DetectLanguage scores close to 1 indicate 100% certainty that the identified language is true. A total of 120
// languages are supported.
// Parameters:
// input - collection of documents to analyze.
func (client BaseClient) DetectLanguage(ctx context.Context, input BatchInput) (result LanguageBatchResult, err error) {
	req, err := client.DetectLanguagePreparer(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "DetectLanguage", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectLanguageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "DetectLanguage", resp, "Failure sending request")
		return
	}

	result, err = client.DetectLanguageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "DetectLanguage", resp, "Failure responding to request")
	}

	return
}

// DetectLanguagePreparer prepares the DetectLanguage request.
func (client BaseClient) DetectLanguagePreparer(ctx context.Context, input BatchInput) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/text/analytics", urlParameters),
		autorest.WithPath("/v2.0/languages"),
		autorest.WithJSON(input))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectLanguageSender sends the DetectLanguage request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DetectLanguageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectLanguageResponder handles the response to the DetectLanguage request. The method always
// closes the http.Response Body.
func (client BaseClient) DetectLanguageResponder(resp *http.Response) (result LanguageBatchResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// KeyPhrases we employ techniques from Microsoft Office's sophisticated Natural Language Processing toolkit. See the
// <a href="https://docs.microsoft.com/en-us/azure/cognitive-services/text-analytics/overview#supported-languages">Text
// Analytics Documentation</a> for details about the languages that are supported by key phrase extraction.
// Parameters:
// input - collection of documents to analyze. Documents can now contain a language field to indicate the text
// language
func (client BaseClient) KeyPhrases(ctx context.Context, input MultiLanguageBatchInput) (result KeyPhraseBatchResult, err error) {
	req, err := client.KeyPhrasesPreparer(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "KeyPhrases", nil, "Failure preparing request")
		return
	}

	resp, err := client.KeyPhrasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "KeyPhrases", resp, "Failure sending request")
		return
	}

	result, err = client.KeyPhrasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "KeyPhrases", resp, "Failure responding to request")
	}

	return
}

// KeyPhrasesPreparer prepares the KeyPhrases request.
func (client BaseClient) KeyPhrasesPreparer(ctx context.Context, input MultiLanguageBatchInput) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/text/analytics", urlParameters),
		autorest.WithPath("/v2.0/keyPhrases"),
		autorest.WithJSON(input))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// KeyPhrasesSender sends the KeyPhrases request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) KeyPhrasesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// KeyPhrasesResponder handles the response to the KeyPhrases request. The method always
// closes the http.Response Body.
func (client BaseClient) KeyPhrasesResponder(resp *http.Response) (result KeyPhraseBatchResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Sentiment scores close to 1 indicate positive sentiment, while scores close to 0 indicate negative sentiment.
// Sentiment score is generated using classification techniques. The input features to the classifier include n-grams,
// features generated from part-of-speech tags, and word embeddings. See the <a
// href="https://docs.microsoft.com/en-us/azure/cognitive-services/text-analytics/overview#supported-languages">Text
// Analytics Documentation</a> for details about the languages that are supported by sentiment analysis.
// Parameters:
// input - collection of documents to analyze.
func (client BaseClient) Sentiment(ctx context.Context, input MultiLanguageBatchInput) (result SentimentBatchResult, err error) {
	req, err := client.SentimentPreparer(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "Sentiment", nil, "Failure preparing request")
		return
	}

	resp, err := client.SentimentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "Sentiment", resp, "Failure sending request")
		return
	}

	result, err = client.SentimentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "textanalytics.BaseClient", "Sentiment", resp, "Failure responding to request")
	}

	return
}

// SentimentPreparer prepares the Sentiment request.
func (client BaseClient) SentimentPreparer(ctx context.Context, input MultiLanguageBatchInput) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/text/analytics", urlParameters),
		autorest.WithPath("/v2.0/sentiment"),
		autorest.WithJSON(input))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SentimentSender sends the Sentiment request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) SentimentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SentimentResponder handles the response to the Sentiment request. The method always
// closes the http.Response Body.
func (client BaseClient) SentimentResponder(resp *http.Response) (result SentimentBatchResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
