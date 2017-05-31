Feature: Comparing steps used and stpes implemented to a step map.

    As a programmer I want to:

    Detect steps that don't match a regex in the mapping file because they could signal what to implement next or a drift in language.

    - Steps: the text of all the steps in a collection of features.

    - Steo map/Step map file: steps.json

    - Step map files: steps.json, steps-unmatched.json, steps-unused.json.

    Scenario: Steps that match the patterns in "steps.json" are left unchanged.

    Scenario: Steps that don't match a pattern in "steps.json" are listed in "steps-unmatched.json".

    Scenario: Steps that aren't used in any scenarios, regardless of filtering, are put in "steps-unused.json".
