Feature: the cucumber runner can parse a basic feature file

Scenario: Parse a basic feature
    Given I have a feature file
    When I parse it
    Then It generates these steps