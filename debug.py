import json

def read_json_file(file):
    with open(file) as f:
        data = json.load(f)
    return data

expected = read_json_file(r'C:\Users\athar\gitleaks\testdata\expected\report\json_with_quotes.json')
actual = read_json_file(r'C:\Users\athar\gitleaks\testdata\expected\test_results\with_quotes.json')

print(expected == actual)