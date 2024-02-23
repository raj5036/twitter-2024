"use client"

import { signIn, signOut, useSession } from "next-auth/react";
import SignInPage from "./components/SignInPage";

export default function Home() {
  const {data: session} = useSession();
  if (session && session.user) {
    console.log(session.user)
    return (
      <div>
        {`User is logged In. Lets show Twitter's feed`}
        <button onClick={() => signOut()}>Sign out</button>
      </div>
    )
  } else {
    return (
      <div>
        <SignInPage />
        <button onClick={() => signIn("google")}>Sign in</button>
      </div>
    )
  }
}
