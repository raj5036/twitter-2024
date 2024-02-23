import { signIn } from "next-auth/react"
import Image from "next/image"

const SignInPage = () => {
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
		  >
			Create account
		  </button>
		  <p className="w-[400px] text-[11px] font-[400] text-twitter-grey">
		  	By signing up, you agree to the  
			<a href="https://twitter.com/en/tos" target="_blank" className="text-twitter-blue"> Terms of Service </a>
			and 
			<a href="https://twitter.com/en/privacy" target="_blank" className="text-twitter-blue"> Privacy Policy </a>
			, including
			<a href="https://help.twitter.com/en/rules-and-policies/x-cookies" target="_blank" className="text-twitter-blue"> Cookie Use</a>.
		  </p>
		  <p>Already have an account?</p>
		  <button
			type="submit"
			className="min-h-[36px] min-w-[380px] rounded-xl border-2 border-solid border-twitter-blue bg-black font-bold text-twitter-blue"
		  >
			Sign in
		  </button>
		</div>
	  </div>
	)
}

export default SignInPage

