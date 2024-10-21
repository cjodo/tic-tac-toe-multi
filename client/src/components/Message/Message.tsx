
interface MessageProps {
		message: string
}

export const Message = ({message}: MessageProps) => {
		return (
		<div>{message}</div>
		)
}
