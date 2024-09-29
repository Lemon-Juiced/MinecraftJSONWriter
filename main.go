package main

import (
	"fmt"
	"os"
)

/**
 * JSONWriter is an application that takes a command line argument containing a namespace, the name of a resource, and a -i, -t, or -c flag denoting an item or tag (either in the given or common directory).
 * It then reads the arguments and writes the data to several JSON files.
 * These JSON files are for Minecraft models and recipes, for a set of tools and armor.
 * The JSON files are written to the correct directories, based on the given namespace for a NeoForge mod.
 *
 * The command line arguments are as follows:
 * make
 * mcjsonwriter <namespace> <resource_name> <-i, -t, or -c>
 *
 * @author Lemon_Juiced & TechDahTurtle
 */
func main() {
	// Check for the correct number of arguments and save them.
	if len(os.Args) != 4 {
		fmt.Println("Usage: mcjsonwriter <namespace> <resource_name> <-i, -t, or -c>")
		return
	}

	namespace := os.Args[1]
	resource_name := os.Args[2]
	//flag := os.Args[3]

	// Create the assets/<namespace>/models/item directory if it doesn't exist and move to it.
	createDirectory("assets/" + namespace + "/models/item")

	// Write the item model JSON files.
	writeItemModelJSONs(namespace, resource_name)
}

/**
 * Create a directory if it doesn't exist and move to it.
 *
 * @param directory The directory to create and move to.
 */
func createDirectory(directory string) {
	// Create the directory if it doesn't exist and move to it.
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.MkdirAll(directory, os.ModePerm)
	}
	os.Chdir(directory)
}
