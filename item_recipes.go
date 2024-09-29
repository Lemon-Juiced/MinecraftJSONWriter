package main

/**
 * Writes the item recipe JSON files for the given resource name and namespace.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeItemRecipeJSONs(namespace string, resource_name string, flag string) {
	// Write the tools JSON files.
	writeAxeRecipeJSON(namespace, resource_name, flag)
	writeHoeRecipeJSON(namespace, resource_name, flag)
	writePickaxeRecipeJSON(namespace, resource_name, flag)
	writeShovelRecipeJSON(namespace, resource_name, flag)
	writeSwordRecipeJSON(namespace, resource_name, flag)

	// Write the armor JSON files.
	writeBootsRecipeJSON(namespace, resource_name, flag)
	writeChestplateRecipeJSON(namespace, resource_name, flag)
	writeHelmetRecipeJSON(namespace, resource_name, flag)
	writeLeggingsRecipeJSON(namespace, resource_name, flag)
}

/**
 * Writes the JSON file for the axe recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeAxeRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"##\",\n" +
		"    \"#\",\n" +
		"    \" /\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    },\n" +
		"    \"/\": {\n" +
		"      \"item\": \"minecraft:stick\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_axe\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_axe.json", content)
}

/**
 * Writes the JSON file for the pickaxe recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writePickaxeRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"###\",\n" +
		"    \" / \",\n" +
		"    \" / \"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    },\n" +
		"    \"/\": {\n" +
		"      \"item\": \"minecraft:stick\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_pickaxe\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_pickaxe.json", content)
}

/**
 * Writes the JSON file for the hoe recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeHoeRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"##\",\n" +
		"    \" /\",\n" +
		"    \" /\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    },\n" +
		"    \"/\": {\n" +
		"      \"item\": \"minecraft:stick\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_hoe\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_hoe.json", content)
}

/**
 * Writes the JSON file for the shovel recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeShovelRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"#\",\n" +
		"    \"/\",\n" +
		"    \"/\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    },\n" +
		"    \"/\": {\n" +
		"      \"item\": \"minecraft:stick\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_shovel\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_shovel.json", content)
}

/**
 * Writes the JSON file for the sword recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeSwordRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"#\",\n" +
		"    \"#\",\n" +
		"    \"/\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    },\n" +
		"    \"/\": {\n" +
		"      \"item\": \"minecraft:stick\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_sword\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_sword.json", content)
}

/**
 * Writes the JSON file for the boots recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeBootsRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"# #\",\n" +
		"    \"# #\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_boots\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_boots.json", content)
}

/**
 * Writes the JSON file for the chestplate recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeChestplateRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"# #\",\n" +
		"    \"###\",\n" +
		"    \"###\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_chestplate\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_chestplate.json", content)
}

/**
 * Writes the JSON file for the helmet recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeHelmetRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"###\",\n" +
		"    \"# #\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_helmet\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_helmet.json", content)
}

/**
 * Writes the JSON file for the leggings recipe.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 */
func writeLeggingsRecipeJSON(namespace string, resource_name string, flag string) {
	content := "{\n" +
		"  \"type\": \"minecraft:crafting_shaped\",\n" +
		"  \"pattern\": [\n" +
		"    \"###\",\n" +
		"    \"# #\",\n" +
		"    \"# #\"\n" +
		"  ],\n" +
		"  \"key\": {\n" +
		"    \"#\": {\n" +
		"      \"" + getKeyType(flag) + "\": \"" + getKeyValue(namespace, resource_name, flag) + "\"\n" +
		"    }\n" +
		"  },\n" +
		"  \"result\": {\n" +
		"    \"item\": \"" + namespace + ":" + resource_name + "_leggings\"\n" +
		"  }\n" +
		"}"

	writeJsonFile(resource_name+"_leggings.json", content)
}

/**
 * Returns the key type based on the flag.
 *
 * @param flag The flag for the item.
 * @return The key type.
 */
func getKeyType(flag string) string {
	if flag == "-i" {
		return "item"
	}
	return "tag"
}

/**
 * Returns the key value based on the namespace, resource name, and flag.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param flag The flag for the item.
 * @return The key value.
 */
func getKeyValue(namespace string, resource_name string, flag string) string {
	if flag == "-c" {
		return "c:" + resource_name
	}
	return namespace + ":" + resource_name
}
