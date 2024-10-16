package pkg

func DefaultPrompt() string {
	return `You will receive some system information along with the most recently run terminal command on the system. You may also receive the output of this command. 
	
	Given the system information and command input + output, dignose the problem and provide a solution. 
	
	Do NOT repeat or summarize the system information.
	
	If there are any obvious typos in the command, provide a suggested command IN THE LAST LINE of your reply, prefixed with "Recommended command: ". The command should be on the same line as "Recommended command: " and there should be no text following it. Do not include any backticks surrounding the command and do not include any lines of text after it.`
}
