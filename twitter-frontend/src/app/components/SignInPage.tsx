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
		  <h1 className="text-[64px] font-extrabold text-[#E7E9EA]">
			Happening Now
		  </h1>
		  <p className="text-[31px] font-extrabold text-[#E7E9EA]">
			Join today
		  </p>
  
		  {/* Google OAuth */}
		  <button
			type="submit"
			className="min-h-[36px] min-w-[380px] rounded-xl border-twitter-blue bg-twitter-blue font-bold text-white"
			onClick={() => signIn('google')}
		  >
			Sign up with Google
		  </button>
  
		  {/* Todo: Work on Apple OAuth */}
		  {/* <button
			type="submit"
			className="min-h-[36px] min-w-[380px] rounded-xl border-twitter-blue bg-twitter-blue font-bold text-white"
		  >
			{t('create_account')}
		  </button> */}
		  <div className="mx-0 mb-[20px] mt-[10px] w-[370px] border-b-[1px] border-b-white bg-black text-center leading-[0.1em]">
			<span className="bg-black px-[10px] py-0 text-white">Or</span>
		  </div>
		  <button
			type="submit"
			className="min-h-[36px] min-w-[380px] rounded-xl border-twitter-blue bg-twitter-blue font-bold text-white"
		  >
			Create account
		  </button>
		  <p className="w-[400px] text-[11px] font-[400] text-twitter-grey">
		  	By signing up, you agree to the Terms of Service and Privacy Policy, including Cookie Use.
		  </p>
		  <p>Already have an account</p>
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

