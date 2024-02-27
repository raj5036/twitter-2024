"use client"

import React, { Fragment, useState } from 'react'
import { default as ReactModal } from "react-modal";

interface ModalProps {
	width: string;
	shouldCloseOnOverlayClick: boolean;
	isOpen: boolean;
	onAfterOpen: () => void;
	onAfterClose: () => void;
	onRequestClose: () => void;
}

const Modal: React.FC<ModalProps> = ({
	isOpen,
	onAfterOpen,
	onAfterClose,
	onRequestClose,
	shouldCloseOnOverlayClick,
	width
}) => {
	// https://reactcommunity.org/react-modal/styles/
	const modalStyles = {
		overlay: {
			backgroundColor: "rgba(0, 0, 0, 0.75)",
		},
		content: {
			top: "50%",
			left: "50%",
			right: "auto",
			bottom: "auto",
			marginRight: "-50%",
			transform: "translate(-50%, -50%)",
			width: width || "auto",
			maxWidth: "900px",
			overflow: "unset",
		},
	};

	return (
		<Fragment>
			<ReactModal
				isOpen={isOpen}
				onAfterOpen={onAfterOpen}
				onAfterClose={onAfterClose}
				onRequestClose={onRequestClose}
				style={modalStyles}
				shouldCloseOnOverlayClick={shouldCloseOnOverlayClick}
				ariaHideApp={false}
			>
				This is a sample Modal
			</ReactModal>
		</Fragment>
	)
}

export default Modal