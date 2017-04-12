package should

import "testing"

//TestJSON exercise the JSON tests
func TestJsonApi(t *testing.T) {
	t.Parallel()
	t.Run("JSON:API record assertions", testBeJSONAPIRecord)
	t.Run("JSON:API resource identifier assertions", testBeResourceIdentifier)
}

func testBeJSONAPIRecord(t *testing.T) {
	json, err := parseJSON(singleRecordResponse)
	Passes(t, "Self documents if passed nil", StartWith, BeJSONAPIRecord(nil), "BeJSONAPIRecord expects")
	Passes(t, "Record parses", BeNil, err)
	Passes(t, "Looks like a JSON:API single-record response.", BeJSONAPIRecord, json)
}

func testBeResourceIdentifier(t *testing.T) {
	json, err := parseJSON(resourceIdentifier)
	Passes(t, "JSON parses", BeNil, err)
	Passes(t, "Looks like a JSON:API resource identifier.", BeJSONAPIResourceIdentifier, json)
	Passes(t, "Self documents if passed nil", StartWith,
		BeJSONAPIResourceIdentifier(nil), "BeJSONAPIResourceIdentifier expects")

	json, _ = parseJSON(resourceIdentifier)
	json.Delete("type")
	Fails(t, "Fails without 'type' field", BeJSONAPIResourceIdentifier, json)
}

const resourceIdentifier = `{
	"id": "abcd123"
	, "type": "person"
}`

const linksSection = `"{
	"self": "http://example.com/articles"
	, "next": "http://example.com/articles?page[offset]=2"
	, "last": "http://example.com/articles?page[offset]=10"
}`

const singleRecordResponse = `{
  "meta": {}
  , "data":
		{ "type": "articles"
    , "id": "1"
    , "attributes": { "title": "JSON API paints my bikeshed!" }
		, "links": { "self": "http://example.com/articles/1" }
		, "relationships":
		  { "author":
				{ "links":
					{ "self": "http://example.com/articles/1/relationships/author"
					, "related": "http://example.com/articles/1/author"
          }
				, "data": { "type": "people", "id": "9" }
        }
			, "comments":
				{ "links":
					{ "self": "http://example.com/articles/1/relationships/comments"
					, "related": "http://example.com/articles/1/comments"
          }
				, "data": [
          { "type": "comments", "id": "5" }
					, { "type": "comments", "id": "12" }
        ]
      }
    }
  }
	, "included": [
		{ "type": "people"
		, "id": "9"
		, "attributes": { "firstName": "Dan" , "lastName": "Gebhardt" , "twitter": "dgeb" }
		, "links": { "self": "http://example.com/people/9" }
    }
	  , { "type": "comments"
	  , "id": "5"
		, "attributes": { "body": "First!" }
		, "relationships": { "author": { "data": { "type": "people", "id": "2" } } }
		, "links": { "self": "http://example.com/comments/5" }
    }
	  , { "type": "comments"
	  , "id": "12"
		, "attributes": { "body": "I like XML better" }
		, "relationships": { "author": { "data": { "type": "people", "id": "9" } } }
		, "links": { "self": "http://example.com/comments/12" }
    }
]
}`

const arrayResponse = `{"meta": {}
, data": [
	{ "type": "articles"
	, "id": "1"
  , "attributes": {
      "title": "JSON API paints my bikeshed!"
    },
    "links": {
      "self": "http://example.com/articles/1"
    },
    "relationships": {
      "author": {
        "links": {
          "self": "http://example.com/articles/1/relationships/author",
          "related": "http://example.com/articles/1/author"
        },
        "data": { "type": "people", "id": "9" }
      },
      "comments": {
        "links": {
          "self": "http://example.com/articles/1/relationships/comments",
          "related": "http://example.com/articles/1/comments"
        },
        "data": [
          { "type": "comments", "id": "5" },
          { "type": "comments", "id": "12" }
        ]
      }
    }
  }],
  "included": [{
    "type": "people",
    "id": "9",
    "attributes": {
      "firstName": "Dan",
      "lastName": "Gebhardt",
      "twitter": "dgeb"
    },
    "links": {
      "self": "http://example.com/people/9"
    }
  }, {
    "type": "comments",
    "id": "5",
    "attributes": {
      "body": "First!"
    },
    "relationships": {
      "author": {
        "data": { "type": "people", "id": "2" }
      }
    },
    "links": {
      "self": "http://example.com/comments/5"
    }
  }, {
    "type": "comments",
    "id": "12",
    "attributes": {
      "body": "I like XML better"
    },
    "relationships": {
      "author": {
        "data": { "type": "people", "id": "9" }
      }
    },
    "links": {
      "self": "http://example.com/comments/12"
    }
  }]
}`

const errorResponse = `{
	"errors":[{
	  "id": "20170101134516-gifts"
		, "links":{
				"about": "https://api.some.com/errors/info/0718"
		  }
		, "status": "400"
		, "code": "0718"
		, "title": "Invalid Filter String"
		, "detail": "'weeek' is not a recognized duration unit"
		, "source": an object containing references to the source of the error, optionally including any of the following members:
		  { "pointer": "/filter/age"
		    , "parameter": "?filter.age=>1weeek,<1year"
			}
		, "meta" {
			apiOrg: "298234785787239487"
	  }
	}],
	"meta": {
	  "method": "GET"
		,"request": "/gifts?filter.age=>1weeek,<1year"
  }
}`
