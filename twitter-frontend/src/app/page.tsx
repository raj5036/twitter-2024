"use client"

import { signIn, signOut, useSession } from "next-auth/react";
import SignInPage from "./components/SignInPage";
import { redirect } from "next/navigation";
import { getUser } from "@/lib/TwitterClient";

export default function Home() {
  const {data: session} = useSession()
  if (session && session.user) {
    // Session user does not exist in Database, then show DOB Modal, save their data & Then redirect to HomePage
    // Session user exists in Database, then redirect to HomePage
    console.log(session.user)
    console.log(getUser({email: session.user.email}))
    // redirect('/home')
  } else {
    return <SignInPage />
  }
}
