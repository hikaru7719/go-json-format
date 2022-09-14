// Generated from /home/hikaru/develop/github.com/hikaru7719/go-json-format/parser/JSON.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JSONLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LEFT_BRACKET=1, RIGHT_BRACKET=2, LEFT_SQUARE_BRACKET=3, RIGHT_SQUARE_BRACKET=4, 
		COMMA=5, COLORN=6, TRUE=7, FALSE=8, NULL=9, DOUBLE_QUOTE=10, WHITE_SPACE=11, 
		STRING=12, INTEGER=13;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET", 
			"COMMA", "COLORN", "TRUE", "FALSE", "NULL", "DOUBLE_QUOTE", "WHITE_SPACE", 
			"STRING", "DIGIT", "INTEGER"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "'['", "']'", "','", "':'", "'true'", "'false'", 
			"'null'", "'\"'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET", 
			"COMMA", "COLORN", "TRUE", "FALSE", "NULL", "DOUBLE_QUOTE", "WHITE_SPACE", 
			"STRING", "INTEGER"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public JSONLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "JSON.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\17a\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3"+
		"\5\3\6\3\6\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n"+
		"\3\n\3\n\3\n\3\13\3\13\3\f\6\f?\n\f\r\f\16\f@\3\f\3\f\3\r\3\r\7\rG\n\r"+
		"\f\r\16\rJ\13\r\3\r\3\r\3\16\3\16\3\17\3\17\3\17\7\17S\n\17\f\17\16\17"+
		"V\13\17\3\17\3\17\3\17\7\17[\n\17\f\17\16\17^\13\17\5\17`\n\17\2\2\20"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\2\35\17\3"+
		"\2\4\5\2\13\f\17\17\"\"\5\2\62;C\\c|\2e\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3"+
		"\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2"+
		"\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\35\3\2\2\2\3\37"+
		"\3\2\2\2\5!\3\2\2\2\7#\3\2\2\2\t%\3\2\2\2\13\'\3\2\2\2\r)\3\2\2\2\17+"+
		"\3\2\2\2\21\60\3\2\2\2\23\66\3\2\2\2\25;\3\2\2\2\27>\3\2\2\2\31D\3\2\2"+
		"\2\33M\3\2\2\2\35_\3\2\2\2\37 \7}\2\2 \4\3\2\2\2!\"\7\177\2\2\"\6\3\2"+
		"\2\2#$\7]\2\2$\b\3\2\2\2%&\7_\2\2&\n\3\2\2\2\'(\7.\2\2(\f\3\2\2\2)*\7"+
		"<\2\2*\16\3\2\2\2+,\7v\2\2,-\7t\2\2-.\7w\2\2./\7g\2\2/\20\3\2\2\2\60\61"+
		"\7h\2\2\61\62\7c\2\2\62\63\7n\2\2\63\64\7u\2\2\64\65\7g\2\2\65\22\3\2"+
		"\2\2\66\67\7p\2\2\678\7w\2\289\7n\2\29:\7n\2\2:\24\3\2\2\2;<\7$\2\2<\26"+
		"\3\2\2\2=?\t\2\2\2>=\3\2\2\2?@\3\2\2\2@>\3\2\2\2@A\3\2\2\2AB\3\2\2\2B"+
		"C\b\f\2\2C\30\3\2\2\2DH\7$\2\2EG\t\3\2\2FE\3\2\2\2GJ\3\2\2\2HF\3\2\2\2"+
		"HI\3\2\2\2IK\3\2\2\2JH\3\2\2\2KL\7$\2\2L\32\3\2\2\2MN\4\62;\2N\34\3\2"+
		"\2\2O`\7\62\2\2PT\4\63;\2QS\5\33\16\2RQ\3\2\2\2SV\3\2\2\2TR\3\2\2\2TU"+
		"\3\2\2\2U`\3\2\2\2VT\3\2\2\2WX\7/\2\2X\\\4\63;\2Y[\5\33\16\2ZY\3\2\2\2"+
		"[^\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]`\3\2\2\2^\\\3\2\2\2_O\3\2\2\2_P\3\2"+
		"\2\2_W\3\2\2\2`\36\3\2\2\2\b\2@HT\\_\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}