"use client"

import { signIn } from "next-auth/react"
// import styles from './SignInPage.module.css';
import { toast } from 'react-toastify';

import Image from "next/image"
import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog"
  import {
	Form,
	FormControl,
	FormDescription,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form"
import {
	Select,
	SelectContent,
	SelectGroup,
	SelectItem,
	SelectLabel,
	SelectTrigger,
	SelectValue,
} from "@/components/ui/select"
import { FloatingLabelInput } from '@/components/ui/floating-label-input';

import {string, z} from 'zod'
import { Button } from "@/components/ui/button"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import axios from 'axios';
import { useEffect } from "react"

const SignInPage = () => {
	const SignInCallbackUrl = '/home'
	const formSchema = z.object({
		name: z.string().min(2, {
			message: "Name must be at least 2 characters.",
		}),
		phone: z.string().min(10, {
			message: "Phone Number must have 10 digits"
		}),
		month: z.string(),
		date: z.string(),
		year: z.string(),
	})

	const form = useForm<z.infer<typeof formSchema>>({
		resolver: zodResolver(formSchema),
		defaultValues: {
			name: "",
			phone: "",
			month: "",
			date: "",
			year: "",
		},
	})

	const onGoogleSignIn = () => {
		signIn("google", {
			callbackUrl: SignInCallbackUrl
		})
	}

	const onAppleSignIn = () => {
		signIn("apple", {
			callbackUrl: SignInCallbackUrl
		})
	}

	const onSubmit = async (data: any) => {
		console.log("onSubmit", Object.keys(data))
		console.log("name", data.name);
		
		if (!data.name || !data.phone || !data.month || !data.date || !data.year) {
			toast("All fields are required!")
			return;
		}
		axios.post('http://localhost:4000/user/create-account', data)
			.then(response => {
				console.log(response)
			})
			.catch(error => {
				console.error("Something went wrong while creating Account", error);
			})
	}

	return (
		<div className="flex flex-row flex-wrap bg-black">
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
					onClick={onGoogleSignIn}
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
					onClick={onAppleSignIn}
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
				<Dialog>
					<DialogTrigger>
						<button
							type="submit"
							className="h-[36px] w-[260px] rounded-[28px] border-twitter-blue bg-twitter-blue font-bold text-white mb-[18px]"
						>
							Create account
						</button>
					</DialogTrigger>
					<CreateAccountDialogContent form={form} onSubmit={onSubmit} />
				</Dialog>
				<p className="w-[400px] text-[11px] font-[400] text-twitter-grey mb-[20px]">
					By signing up, you agree to the  
					<ExternalLink textToDisplay="Terms of Service" url="https://twitter.com/en/tos"/>
					and 
					<ExternalLink textToDisplay="Privacy Policy" url="https://twitter.com/en/privacy"/>
					, including
					<ExternalLink textToDisplay="Cookie Use" url="https://help.twitter.com/en/rules-and-policies/x-cookies"/>
				</p>
				<p
					className="text-white font-bold mt-[40px] mb-[20px]"
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
	  </div>
	)
}

const ExternalLink = ({textToDisplay, url}: {
	textToDisplay: string,
	url: string
}) => {
	return <a href={url} target="_blank" className="text-twitter-blue">{" "}{textToDisplay}{" "}</a>
}

const SelectScrollable = ({placeholder, selectWidth, options, form, field}: {
	placeholder: string,
	selectWidth: string,
	options: Array<string>
	form: any
	field: any
}) => {
	return (
		<Select onValueChange={field.onChange} defaultValue={field.value} key={selectWidth}>
			<SelectTrigger className={`bg-black text-twitter-foreground w-[${selectWidth}]`}>
				<SelectValue placeholder={placeholder}/>
			</SelectTrigger>
			<SelectContent className="bg-black text-twitter-foreground">
				{options.map((option, index) => <SelectItem key={index} value={option}>{option}</SelectItem>)}
			</SelectContent>
		</Select>
	)
}

const CreateAccountDialogContent = ({form, onSubmit}: {
	form: any,
	onSubmit: Function
}) => {
	const Months = Array.from({ length: 12 }).map((_, index) => {
		const monthName = new Date(2000, index).toLocaleDateString('en-US', { month: 'long' });
		return monthName;
	})
	
	const getYearsFrom1904 = () => {
		const currentYear = new Date().getFullYear();
		const years = [];
		
		for (let year = currentYear; year >= 1904; year--) {
			years.push(String(year));
		}
		
		return years;
	}

	const Years = getYearsFrom1904()
	const Dates = Array.from({length: 31}).map((_, index) => (index + 1).toString())

	return <DialogContent className="bg-black text-white">
		<DialogHeader>
			<DialogTitle>
				<Image
					src="/assets/images/twitter-x-logo.png"
					alt="twitter svg"
					width={32}
					height={32}
					className="m-auto"
				/>
			</DialogTitle>
			<DialogDescription>
				<h1 className="text-3xl text-twitter-foreground my-[20px]">Create your account</h1>
				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
					<FormField
						control={form.control}
						name="name"
						render={({ field }) => (
							<FormItem>
								<FloatingLabelInput 
									id="floating-demo"
									label="Name" 
									className="text-white bg-black p-[10px]"
									{...field}
								/>
								<FormMessage />
							</FormItem>
						)}
					/>
					<FormField
						control={form.control}
						name="phone"
						render={({ field }) => (
							<FormItem>
								<FloatingLabelInput 
									id="floating-demo"
									label="Phone" 
									className="text-white bg-black p-[10px]"
									{...field}
								/>
								<FormMessage />
							</FormItem>
						)}
					/>
					<div>
						<h2 className="text-twitter-foreground scroll-m-20 pb-2 text-base font-semibold tracking-tight first:mt-0">
							Date of Birth
						</h2>
						<p className="m-0">
							This will not be shown publicly. Confirm your own age, even if this account is for a business, 
							a pet, or something else.
						</p>
						<div className="flex justify-evenly items-center mt-[2rem]">
							<FormField
								control={form.control}
								name="month"
								render={({ field }) => (
									<FormItem>
										<SelectScrollable placeholder="Month" selectWidth="200px" options={Months} form={form} field={field}/>
									</FormItem>
								)}
							/>
							<FormField
								control={form.control}
								name="date"
								render={({ field }) => (
									<FormItem>
										<SelectScrollable placeholder="Date" selectWidth="100px" options={Dates} form={form} field={field}/>
									</FormItem>
								)}
							/>
							<FormField
								control={form.control}
								name="year"
								render={({ field }) => (
									<FormItem>
										<SelectScrollable placeholder="Year" selectWidth="100px" options={Years} form={form} field={field}/>
									</FormItem>
								)}
							/>
						</div>
					</div>
					<div className="flex justify-center items-center mt-[2rem]">
						<Button 
							type="submit"
							className="w-[28rem] hover:bg-white hover:text-black"
						>Submit</Button>
					</div>
				</form>
				</Form>
			</DialogDescription>
		</DialogHeader>
	</DialogContent>
	
}	

export default SignInPage

