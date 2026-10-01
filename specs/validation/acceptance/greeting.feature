Feature: Greeting

  @story-1
  Rule: A caller who supplies a name receives a greeting that includes it

    Scenario: Greeting a named caller
      Given the greeter service is running
      When an API Consumer calls "GET /hello?name=Ada"
      Then the response is a JSON greeting that includes "Ada"

  @story-2
  Rule: A caller who omits the name still receives a valid default greeting

    Scenario: Greeting with no name supplied
      Given the greeter service is running
      When an API Consumer calls "GET /hello" with no name parameter
      Then the response is a JSON greeting with a default name

    Scenario: Greeting with an empty name supplied
      Given the greeter service is running
      When an API Consumer calls "GET /hello?name="
      Then the response is a JSON greeting with a default name
