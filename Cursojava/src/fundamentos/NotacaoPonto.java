package fundamentos;

public class NotacaoPonto {
	public static void main(String[] args) {
		String s = "Bom dia?".toUpperCase().replace("?", " bia").concat("!");
		System.out.println(s);
	}
}