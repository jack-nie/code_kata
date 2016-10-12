package linecount;

import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Assume code is compilable. (Comments and quotes terminate properly.)
 *
 */
public class CountLines {
	private boolean insideBlockComment = false;
	boolean insideQuote = false;

	public static void main(String[] args) {
		// TODO Auto-generated method stub

	}

	public int countLines(List<String> input) {
		int count = 0;
		for (String line : input) {
			if (isValid(line)) {
				count++;
			}
		}
		return count;
	}

	private boolean isValid(String line) {
		boolean valid = false;

		for (int index = 0; index < line.length(); index++) {
			char ch = line.charAt(index);
			switch (ch) {
				case ' ':
				case '\n':
				case '\t':
				case '\r':
					break;
				case '"':
					if (!insideBlockComment) {
						valid = true;
						insideQuote = !insideQuote;
					}
					break;
				case '/':
					if (isLineComment(line, index) ) {
						if (!insideBlockComment) {
							return valid;
						}
					}

					if (isBlockComment(line, index) ) {
						insideBlockComment = true;
						index++;
					} else {
						if (!insideBlockComment) {
							valid = true;
						}
					}

					break;
				case '*':
					if (insideBlockComment) {
						if (matchedCharAtIndex('/', line, index + 1)) {
							insideBlockComment = false;
							index++;
						}
					} else {
						valid = true;
					}
					break;
				default:
					if (!insideBlockComment) {
						valid = true;
					}
			}
		}
		return valid;
	}

	private boolean isBlockComment(String line, int index) {
		if (matchedCharAtIndex('*', line, index + 1) ) {
			if (!insideQuote) {
				return insideBlockComment = true;
			}
		}
		return false;
	}

	private boolean isLineComment(String line, int index) {
		return (matchedCharAtIndex('/', line, index + 1));
	}

	private boolean matchedCharAtIndex(char searchTarget, String line, int index) {
		if (index == line.length()) {
			return false;
		}
		return (searchTarget == line.charAt(index));
	}

}
