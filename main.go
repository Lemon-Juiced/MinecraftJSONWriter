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
	if len(os.Args) < 2 {
		fmt.Println("Please provide arguments.")
		return
	}

	for i, arg := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", i+1, arg)
	}
}
