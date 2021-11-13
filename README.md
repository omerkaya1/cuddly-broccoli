# cuddly-broccoli - Convert JSON objects to an array of objects

A package that aims to solve one agonising problem - convert a poorly formed JSON with objects that contain a collection
of identical objects that should have been an array of objects.
It does not rely on reflection, although it does rely on external packages to parse JSON data.

## TL;DR: 
Someone on a team had an ingenious idea to serialise a large (~18k lines of prettified) JSON as ```map[string]interface{}```.
So, the data structure changes every 5-6 months and there is no way they can provide a reasonable schema for it.
Thus, here we are. Introducing new bugs due to bad design decisions.

## API

The package is pretty straightforward:
1. Define paths, using the package conventions
2. Serialise the original data
3. Obtain the converted result
```cgo
Path composing conventions:
    ?# - object keys are stringified numbers and we need to include them to converted objects
    ?* - object keys are names and we need to include them to converted objects
    ?#- - numbers are irrelevant and we can ditch them
    ?*- - names are irrelevant and we can ditch them
Example:
    "example#.path*-.to*.object"
```

## Example:

```json
{
	"show": "Rick and Morty",
	"characters": {
		"0": {
			"name": "Rick",
			"surname": "Sanchez",
			"episodes": {
				"0": 10,
				"1": 12
			},
			"clothes": {
				"0": {"type": "Lab coat", "color": "white", "size": "M"},
				"1": {"type": "Pents", "color": "brown", "size": "M"},
				"2": {"type": "T-Short", "color": "brown", "size": "M"}
			}
		},
		"1": {
			"name": "Morty",
			"surname": "Smith",
			"episodes": {
				"0": 10,
				"1": 12
			},
			"clothes": {
				"0": {"type": "Pents", "color": "brown", "size": "S"},
				"1": {"type": "T-Short", "color": "brown", "size": "S"}
			}
		}
	},
	"episodes": {
		"1": {
			"name": "Pilot",
			"characters": {
				"0": {"name": "Rick"},
				"1": {"name": "Morty"},
				"2": {"name": "Summer"},
				"3": {"name": "Beth"},
				"4": {"name": "Jerry"}
			}
		},
		"2": {
			"name": "Second",
			"characters": {
				"0": {"name": "Rick"},
				"1": {"name": "Morty"},
				"2": {"name": "Summer"},
				"3": {"name": "Beth"},
				"4": {"name": "Jerry"}
			}
		}
	}
}
```

Clearly, the structure above could have been much simpler were it composed of arrays instead of objects.

So, the project aims at rectifying this issue by parsing the serialised object and convert it to an idiomatic JSON 
object based on the "paths" the user provides.
Given that the paths are composed in the following manner:
```cgo
var rickAndMortyPaths = []string{
    "characters#-.clothes#-",
    "characters#-.episodes#-",
    "episodes#.characters#-",
}
```
With paths composed the way outlined above, the above JSON becomes:
```json
{
	"show": "Rick and Morty",
	"characters": [
		{
			"name": "Rick",
			"surname": "Sanchez",
			"clothes": [
				{"type": "Lab coat", "color": "white", "size": "M"},
				{"type": "Pents", "color": "brown", "size": "M"},
				{"type": "T-Short", "color": "brown", "size": "M"}
			]
		},
		{
			"name": "Morty",
			"surname": "Smith",
			"clothes": [
				{"type": "Pents", "color": "brown", "size": "S"},
				{"type": "T-Short", "color": "yelow", "size": "S"}
			]
		}
	],
	"episodes": [
		{
			"name": "Pilot",
			"characters": [
				{"name": "Rick"},
				{"name": "Morty"},
				{"name": "Summer"},
				{"name": "Beth"},
				{"name": "Jerry"}
			]
		},
		{
			"name": "Second",
			"characters": [
				{"name": "Rick"},
				{"name": "Morty"},
				{"name": "Summer"},
				{"name": "Beth"},
				{"name": "Jerry"}
			]
		}
	]
}
```
