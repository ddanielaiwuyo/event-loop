#!/usr/bin/env node
"use strict"

// pending_promise <- new Promise(func(res, rej))
// pending_promise.then((success), (failure)) // basically only 
// one could be true at the same time? once success fires, thats it
function krang() {
	// this part is syncrhonouse as we haven't called it yet, just declared
	let promise = new Promise(function (resolve, reject) {
		setTimeout(() => resolve("weird_fishes"), 2000)
		setTimeout(() => reject(new Error("whoopsy_daisy")), 2000)
	})

	console.log("you are not the one the krang expected, not expected are you the krang")

	// I think this should block?
	let new_v = null
	// this part is popped into the promise queue
	promise.then((v) => {
		console.log(v)
		new_v = v
	}, (err) => console.log(err))

	console.log(new_v)

}

import fs from "node:fs/promises"
async function reader(file_name) {
	try {
		// popped to ioq
		const fh = await fs.open(file_name, "r")
		// ran immediately
		console.log("this should not run until the content has been resolved")

		// popped into queue, but its dependent on fh
		for await (const line of fh.readLines()) {
			console.log(line)
		}

		console.log("caravan")

	} catch (err) {
		console.error(err)
	}
}

await reader("io.go")
console.log("runtime/mock.js")
