"use client"

import { signIn, signOut, useSession } from "next-auth/react"

const SignInButton = () => {
	const {data: session} = useSession();
	if (session && session.user) {
		console.log("userSession", session.user)
		return (<>
			<h1>User already signed in</h1>
			<button onClick={() => signOut()}>Sign out</button>
		</>
		)
	}
	return (
		<button onClick={() => signIn("google")}>Sign in</button>
	)
}

export default SignInButton