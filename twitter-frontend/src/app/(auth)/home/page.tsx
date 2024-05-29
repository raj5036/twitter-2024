'use client'
import { useSession } from "next-auth/react"
import React, {useState, useEffect} from "react"

const Page = () => {
	const { data: session } = useSession()
	console.log('session', session)

	// Check if email exists in Database, else redirect to SignIn page
	return (
		<h1>home page</h1>
	)
}

export default Page