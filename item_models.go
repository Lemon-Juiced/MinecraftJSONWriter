package main

/**
 * Writes the item model JSON files for the given resource name and namespace.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 */
func writeItemModelJSONs(namespace string, resource_name string) {
	// Write the tools JSON files.
	writeToolModelJSON(namespace, resource_name, "axe")
	writeToolModelJSON(namespace, resource_name, "hoe")
	writeToolModelJSON(namespace, resource_name, "pickaxe")
	writeToolModelJSON(namespace, resource_name, "shovel")
	writeToolModelJSON(namespace, resource_name, "sword")

	// Write the armor JSON files.
	writeArmorModelJSON(namespace, resource_name, "boots")
	writeArmorModelJSON(namespace, resource_name, "chestplate")
	writeArmorModelJSON(namespace, resource_name, "helmet")
	writeArmorModelJSON(namespace, resource_name, "leggings")
}

/**
 * Write the tools JSON files for the given resource name and namespace.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param tool_type The type of tool.
 */
func writeToolModelJSON(namespace string, resource_name string, tool_type string) {
	// Write the tool item model JSON file.
	content := "{\n" +
		"  \"parent\": \"item/handheld\",\n" +
		"  \"textures\": {\n" +
		"    \"layer0\": \"" + namespace + ":items/" + resource_name + "_" + tool_type + "\"\n" +
		"  }\n" +
		"}"
	writeJsonFile(resource_name+"_"+tool_type+".json", content)
}

/**
 * Write the armor JSON files for the given resource name and namespace.
 * Additionally write the trimmed armor JSON files for the given resource name and namespace.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param armor_type The type of armor.
 */
func writeArmorModelJSON(namespace string, resource_name string, armor_type string) {
	// Write the armor item model JSON file with overrides for trimmed variants.
	content := "{\n" +
		"  \"parent\": \"minecraft:item/generated\",\n" +
		"  \"overrides\": [\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_quartz_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.1\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_iron_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.2\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_netherite_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.3\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_redstone_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.4\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_copper_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.5\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_gold_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.6\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_emerald_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.7\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_diamond_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.8\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_lapis_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 0.9\n" +
		"      }\n" +
		"    },\n" +
		"    {\n" +
		"      \"model\": \"" + namespace + ":item/" + resource_name + "_" + armor_type + "_amethyst_trim\",\n" +
		"      \"predicate\": {\n" +
		"        \"trim_type\": 1.0\n" +
		"      }\n" +
		"    }\n" +
		"  ],\n" +
		"  \"textures\": {\n" +
		"    \"layer0\": \"" + namespace + ":items/" + resource_name + "_" + armor_type + "\"\n" +
		"  }\n" +
		"}"
	writeJsonFile(resource_name+"_"+armor_type+".json", content)

	// Write the trimmed armor item model JSON files.
	trims := []string{"amethyst", "copper", "diamond", "emerald", "gold", "iron", "lapis", "netherite", "quartz", "redstone"}
	for _, trim := range trims {
		trimContent := "{\n" +
			"  \"parent\": \"minecraft:item/generated\",\n" +
			"  \"textures\": {\n" +
			"    \"layer0\": \"" + namespace + ":items/" + resource_name + "_" + armor_type + "\",\n" +
			"    \"layer1\": \"" + namespace + ":trims/items/" + armor_type + "_trim_" + trim + "\"\n" +
			"  }\n" +
			"}"
		writeJsonFile(resource_name+"_"+armor_type+"_"+trim+"_trim.json", trimContent)
	}
}
