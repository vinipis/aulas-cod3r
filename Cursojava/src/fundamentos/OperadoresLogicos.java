package fundamentos;

public class OperadoresLogicos {

	public static void main(String[] args) {
		boolean executouTrab1 = false;
		boolean executouTrab2 = true;
		
		boolean comprouSorvete = executouTrab1 || executouTrab2;
		boolean comprouTv50 = executouTrab1 && executouTrab2;
		boolean comprouTv32 = executouTrab1 ^ executouTrab2;
		
		System.out.println("Sorvete " + comprouSorvete);
		System.out.println("TV50 " + comprouTv50);
		System.out.println("TV32 " + comprouTv32);
		
		//Operador unario intruso
		System.out.println("Fome = " + !comprouSorvete);
	}
}
