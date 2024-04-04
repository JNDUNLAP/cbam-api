import json

def convert_pattern_to_go_regex(pattern):
    """
    Converts XML Schema regex patterns to Go's regex format.
    Basic conversion; more complex patterns might require advanced handling.
    """
    return pattern.replace("\\d", "[0-9]")

def generate_go_struct(json_input_file, go_output_file):
    with open(json_input_file, 'r') as json_file:
        data = json.load(json_file)

    needs_regexp_import = any("pattern" in details.get("restrictions", {}) for details in data.values())
    import_statements = 'import "time"\n'
    # if needs_regexp_import:
    #     import_statements += 'import "regexp"\n'
    
    go_struct = [f"package model\n\n{import_statements}\n// Report represents the structure of our report\ntype Report struct {{"]

    for field, details in data.items():
        go_type = "string"  # Default to string
        validation_tags = []

        if details["type"].startswith("DateTimeContentType"):
            go_type = "time.Time"
        elif details["type"].startswith("an") or details["type"].startswith("n"):
            go_type = "string"
        
        if "minLength" in details.get("restrictions", {}):
            validation_tags.append(f'min={details["restrictions"]["minLength"]}')
        if "maxLength" in details.get("restrictions", {}):
            validation_tags.append(f'max={details["restrictions"]["maxLength"]}')
        if "totalDigits" in details.get("restrictions", {}):
            validation_tags.append(f'max={10**int(details["restrictions"]["totalDigits"]) - 1}')
        if "pattern" in details.get("restrictions", {}):
            go_regex = convert_pattern_to_go_regex(details["restrictions"]["pattern"])
            validation_tags.append(f'regex=\"{go_regex}\"')

        go_field_name = ''.join(word.capitalize() for word in field.split('_'))
        tag_string = ','.join(validation_tags)
        go_struct.append(f'\t{go_field_name} {go_type} `json:"{field}" validate:"{tag_string}"`')

    go_struct.append("}")  # Close the struct definition

    with open(go_output_file, 'w') as go_file:
        go_file.write('\n'.join(go_struct))

    print(f"Go struct has been generated in {go_output_file}")
