package fundamentos;

public class ConversaoNumerica {

	public static void main(String[] args) {
		//Exemplo 1 - Conversão explicita
		float f = (float) 0.1;
		System.out.println(f);
		
		//exemplo 2 - conversão explicita
		int i1 = (int) 7.9;
		System.out.println(i1);
		
		//exemplo 3 - conversão implicita
		int i2 = 'a';
		System.out.println(i2);
		
		//exemplo 4 - conversão implicita
		double d = 1001;
		System.out.println(d);
	}
}
