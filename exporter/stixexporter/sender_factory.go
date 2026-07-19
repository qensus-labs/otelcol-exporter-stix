package stixexporter

func createSender(
	output string,
) (Sender, error) {

	switch output {

	case "", "stdout":
		return newStdoutSender(), nil

	default:

		sender, err := newFileSender(
			output,
		)

		if err != nil {
			return nil, err
		}

		return sender, nil
	}
}
