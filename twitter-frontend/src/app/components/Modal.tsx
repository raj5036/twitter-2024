"use client"

import React, { Fragment, useState } from 'react'

interface ModalProps extends React.HTMLAttributes<any>{
	width: string;
	shouldCloseOnOverlayClick: boolean;
	isOpen: boolean;
	onAfterOpen?: () => void;
	onAfterClose?: () => void;
	onRequestClose: () => void;
	children: React.ReactNode;
}

const Modal: React.FC<ModalProps> = ({
	isOpen,
	onRequestClose,
	children
}) => {
	if (!isOpen) return null;
	return (
		<Fragment>
			{children}
			<div>
				<button onClick={onRequestClose}>
					Close
				</button>
			</div>
		</Fragment>
	)
}

export default Modal