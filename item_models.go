package main

/**
 * Writes the item model JSON files for the given resource name and namespace.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 */
func writeItemModelJSONs(namespace string, resource_name string) {
	// Write the tools JSON files.
	writeToolModelJson(namespace, resource_name, "axe")
	writeToolModelJson(namespace, resource_name, "hoe")
	writeToolModelJson(namespace, resource_name, "pickaxe")
	writeToolModelJson(namespace, resource_name, "shovel")
	writeToolModelJson(namespace, resource_name, "sword")

	// Write the armor JSON files.
	writeArmorModelJson(namespace, resource_name, "boots")
	writeArmorModelJson(namespace, resource_name, "chestplate")
	writeArmorModelJson(namespace, resource_name, "helmet")
	writeArmorModelJson(namespace, resource_name, "leggings")
}

/**
 * Write the tools JSON files for the given resource name and namespace.
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param tool_type The type of tool.
 */
func writeToolModelJson(namespace string, resource_name string, tool_type string) {
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
 *
 * @param namespace The namespace for the mod.
 * @param resource_name The name of the resource.
 * @param armor_type The type of armor.
 */
func writeArmorModelJson(namespace string, resource_name string, armor_type string) {
	// Write the armor item model JSON file.
	content := "{\n" +
		"  \"parent\": \"item/generated\",\n" +
		"  \"textures\": {\n" +
		"    \"layer0\": \"" + namespace + ":items/" + resource_name + "_" + armor_type + "\"\n" +
		"  }\n" +
		"}"
	writeJsonFile(resource_name+"_"+armor_type+".json", content)
}
