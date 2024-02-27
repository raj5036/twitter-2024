"use client"

import { signIn } from "next-auth/react"

import Image from "next/image"
import { useState } from "react"
import Modal from "./Modal"

const SignInPage = () => {
	const [createAccountModalOpen, setCreateAccountModal] = useState<boolean>(false);
	const openCreateAccountModal = () => {
		console.log("openCreateAccountModal")
		setCreateAccountModal(true);
	}

	return (
		<div className="flex h-[100vh] flex-row flex-wrap bg-black">
		<div className="h-[100%] w-[50%]">
		  <Image
			src="/assets/images/twitter-x-logo.png"
			alt="twitter svg"
			width={400}
			height={260}
			className="mx-auto my-[7rem]"
		  />
		</div>
		<div className="h-[100%] w-[50%] text-left">
			<h1 className="text-[64px] font-extrabold text-[#E7E9EA] mx-0 my-[48px]">
				Happening Now
		  	</h1>
			<p className="text-[31px] font-extrabold text-[#E7E9EA] mb-[32px]">
				Join today
			</p>
  
		  	{/* Google OAuth */}
		  	<div 
		  		className="text-black font-medium text-sm bg-white flex h-[44px] max-w-[232px] py-[13px] px-[35px] rounded-[28px] cursor-pointer mb-[15px]"
				onClick={() => signIn("google")}
		   	>
				<Image
					src={"/assets/images/google_icon.png"}
					alt="Google Icon"
					height={20}
					width={20}
				/>
				<span>Sign up with Google</span>
		  	</div>
  
		  	{/* Todo: Work on Apple OAuth */}
		  	<div 
		  		className="text-black font-bold text-sm bg-white flex h-[44px] max-w-[232px] py-[13px] px-[35px] rounded-[28px] cursor-pointer"
				onClick={() => signIn("apple")}
		   	>
				<Image
					src={"/assets/images/apple-logo.png"}
					alt="Google Icon"
					height={20}
					width={30}
				/>
				<span>Sign up with Apple</span>
			</div>
			<div className="mx-0 mb-[20px] mt-[30px] w-[255px] border-b-[0.5px] border-b-twitter-grey text-center leading-[0.1em]">
				<span className="bg-black px-[10px] py-0 text-white">or</span>
			</div>
			<button
				type="submit"
				className="h-[36px] w-[260px] rounded-[28px] border-twitter-blue bg-twitter-blue font-bold text-white mb-[18px]"
				onClick={openCreateAccountModal}
			>
				Create account
			</button>
			<p className="w-[400px] text-[11px] font-[400] text-twitter-grey mb-[20px]">
				By signing up, you agree to the  
				<ExternalLink textToDisplay="Terms of Service" url="https://twitter.com/en/tos"/>
				and 
				<ExternalLink textToDisplay="Privacy Policy" url="https://twitter.com/en/privacy"/>
				, including
				<ExternalLink textToDisplay="Cookie Use" url="https://help.twitter.com/en/rules-and-policies/x-cookies"/>
			</p>
			<p
				className="text-twitter-white font-normal mt-[40px] mb-[20px]"
			>
				Already have an account?
			</p>
			<button
				type="submit"
				className="h-[36px] w-[260px] rounded-[28px] border-twitter-grey-2 border-[1px] bg-black font-bold text-twitter-blue mb-[18px]"
			>
				Sign in
			</button>
		</div>
		<Modal
			isOpen={createAccountModalOpen}
			onAfterClose={()=> {}}
			onAfterOpen={()=>{}}
			onRequestClose={()=>{}}
			shouldCloseOnOverlayClick={true}
			width="100px"
			key={1}
		>
			This is a modal
		</Modal>
	  </div>
	)
}

const ExternalLink = ({textToDisplay, url}: {
	textToDisplay: string,
	url: string
}) => {
	return <a href={url} target="_blank" className="text-twitter-blue">{" "}{textToDisplay}{" "}</a>
}

export default SignInPage

