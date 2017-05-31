Feature: Default and alterable configuration.

Scenario: There is a typical file layout.
    Given a base directory ".../base".
    When I use default settings.
    Then The step maps are in the ".../base/".
    And features are expected in ".../base/features/".
    And test files will be written in ".../base/tests/".




