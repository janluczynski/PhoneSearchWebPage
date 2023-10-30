import pandas as pd
import json

# Read the DataFrame from the URL
url = 'https://query.data.world/s/hf2onxwypja2gu56igou3d7vrqltsi?dws=00000'
df = pd.read_csv(url)

# Clean up the 'description' field
df['description'] = df['description'].replace({'\n': ' ', '\r': ' ', '\u201d': '"', '\u201c': '"'}, regex=True)

# Convert DataFrame to a list of dictionaries
data_as_dict_list = df.to_dict(orient='records')

# Save the data as a JSON file
with open('output.json', 'w', encoding='utf-8') as json_file:
    json.dump(data_as_dict_list, json_file, ensure_ascii=False)

print("Data has been saved as 'output.json' in the form of an array of dictionaries with cleaned 'description' fields.")
print("To get a cleaner looking json file use https://jsonformatter.curiousconcept.com/")
