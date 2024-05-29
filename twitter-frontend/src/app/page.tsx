"use client"

import { signIn, signOut, useSession } from "next-auth/react";
import SignInPage from "./components/SignInPage";

export default function Home() {
  const {data: session} = useSession();
  if (session && session.user) {
    /*
      This if statement might be unnecessary. Check if it can be removed.
    */
    console.log(session.user)
    return (
      <div>
        {`User is logged In. Lets show Twitter's feed`}
        <button onClick={() => signOut()}>Sign out</button>
      </div>
    )
  } else {
    return <SignInPage />
  }
}
