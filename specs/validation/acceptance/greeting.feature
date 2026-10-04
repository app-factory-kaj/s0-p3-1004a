Feature: Greeting

  @story-1
  Rule: Requesting a greeting with a name returns a personalized JSON greeting

    Scenario: A client asks for a greeting by name
      Given the greeter service is running
      When a client application sends a GET request to "/hello" with name "Ada"
      Then the response is a JSON body with a message greeting "Ada"

  @story-2
  Rule: Requesting a greeting without a name returns a default JSON greeting

    Scenario: A client omits the name parameter
      Given the greeter service is running
      When a client application sends a GET request to "/hello" with no name parameter
      Then the response is a JSON body with a default greeting message
