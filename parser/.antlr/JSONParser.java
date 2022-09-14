// Generated from /home/hikaru/develop/github.com/hikaru7719/go-json-format/parser/JSON.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JSONParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LEFT_BRACKET=1, RIGHT_BRACKET=2, LEFT_SQUARE_BRACKET=3, RIGHT_SQUARE_BRACKET=4, 
		COMMA=5, COLORN=6, TRUE=7, FALSE=8, NULL=9, DOUBLE_QUOTE=10, WHITE_SPACE=11, 
		STRING=12, INTEGER=13;
	public static final int
		RULE_json = 0, RULE_value = 1, RULE_object = 2, RULE_members = 3, RULE_member = 4, 
		RULE_array = 5, RULE_elements = 6, RULE_str = 7, RULE_num = 8, RULE_bool = 9, 
		RULE_nil = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"json", "value", "object", "members", "member", "array", "elements", 
			"str", "num", "bool", "nil"
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

	@Override
	public String getGrammarFileName() { return "JSON.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public JSONParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class JsonContext extends ParserRuleContext {
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public JsonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_json; }
	}

	public final JsonContext json() throws RecognitionException {
		JsonContext _localctx = new JsonContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_json);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public ObjectContext object() {
			return getRuleContext(ObjectContext.class,0);
		}
		public ArrayContext array() {
			return getRuleContext(ArrayContext.class,0);
		}
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public StrContext str() {
			return getRuleContext(StrContext.class,0);
		}
		public BoolContext bool() {
			return getRuleContext(BoolContext.class,0);
		}
		public NilContext nil() {
			return getRuleContext(NilContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_value);
		try {
			setState(30);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LEFT_BRACKET:
				enterOuterAlt(_localctx, 1);
				{
				setState(24);
				object();
				}
				break;
			case LEFT_SQUARE_BRACKET:
				enterOuterAlt(_localctx, 2);
				{
				setState(25);
				array();
				}
				break;
			case INTEGER:
				enterOuterAlt(_localctx, 3);
				{
				setState(26);
				num();
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 4);
				{
				setState(27);
				str();
				}
				break;
			case TRUE:
			case FALSE:
				enterOuterAlt(_localctx, 5);
				{
				setState(28);
				bool();
				}
				break;
			case NULL:
				enterOuterAlt(_localctx, 6);
				{
				setState(29);
				nil();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectContext extends ParserRuleContext {
		public TerminalNode LEFT_BRACKET() { return getToken(JSONParser.LEFT_BRACKET, 0); }
		public TerminalNode RIGHT_BRACKET() { return getToken(JSONParser.RIGHT_BRACKET, 0); }
		public MembersContext members() {
			return getRuleContext(MembersContext.class,0);
		}
		public ObjectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object; }
	}

	public final ObjectContext object() throws RecognitionException {
		ObjectContext _localctx = new ObjectContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_object);
		try {
			setState(38);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(32);
				match(LEFT_BRACKET);
				setState(33);
				match(RIGHT_BRACKET);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				match(LEFT_BRACKET);
				setState(35);
				members();
				setState(36);
				match(RIGHT_BRACKET);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MembersContext extends ParserRuleContext {
		public List<MemberContext> member() {
			return getRuleContexts(MemberContext.class);
		}
		public MemberContext member(int i) {
			return getRuleContext(MemberContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(JSONParser.COMMA, 0); }
		public MembersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_members; }
	}

	public final MembersContext members() throws RecognitionException {
		MembersContext _localctx = new MembersContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_members);
		try {
			setState(45);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				member();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(41);
				member();
				setState(42);
				match(COMMA);
				setState(43);
				member();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MemberContext extends ParserRuleContext {
		public StrContext str() {
			return getRuleContext(StrContext.class,0);
		}
		public TerminalNode COLORN() { return getToken(JSONParser.COLORN, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public MemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_member; }
	}

	public final MemberContext member() throws RecognitionException {
		MemberContext _localctx = new MemberContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_member);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			str();
			setState(48);
			match(COLORN);
			setState(49);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrayContext extends ParserRuleContext {
		public TerminalNode LEFT_SQUARE_BRACKET() { return getToken(JSONParser.LEFT_SQUARE_BRACKET, 0); }
		public TerminalNode RIGHT_SQUARE_BRACKET() { return getToken(JSONParser.RIGHT_SQUARE_BRACKET, 0); }
		public ElementsContext elements() {
			return getRuleContext(ElementsContext.class,0);
		}
		public ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array; }
	}

	public final ArrayContext array() throws RecognitionException {
		ArrayContext _localctx = new ArrayContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_array);
		try {
			setState(57);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(51);
				match(LEFT_SQUARE_BRACKET);
				setState(52);
				match(RIGHT_SQUARE_BRACKET);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(53);
				match(LEFT_SQUARE_BRACKET);
				setState(54);
				elements();
				setState(55);
				match(RIGHT_SQUARE_BRACKET);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ElementsContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(JSONParser.COMMA, 0); }
		public ElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elements; }
	}

	public final ElementsContext elements() throws RecognitionException {
		ElementsContext _localctx = new ElementsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_elements);
		try {
			setState(64);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				value();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				value();
				setState(61);
				match(COMMA);
				setState(62);
				value();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StrContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(JSONParser.STRING, 0); }
		public StrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_str; }
	}

	public final StrContext str() throws RecognitionException {
		StrContext _localctx = new StrContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_str);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NumContext extends ParserRuleContext {
		public TerminalNode INTEGER() { return getToken(JSONParser.INTEGER, 0); }
		public NumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_num; }
	}

	public final NumContext num() throws RecognitionException {
		NumContext _localctx = new NumContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_num);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(INTEGER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BoolContext extends ParserRuleContext {
		public TerminalNode TRUE() { return getToken(JSONParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(JSONParser.FALSE, 0); }
		public BoolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool; }
	}

	public final BoolContext bool() throws RecognitionException {
		BoolContext _localctx = new BoolContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_bool);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			_la = _input.LA(1);
			if ( !(_la==TRUE || _la==FALSE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NilContext extends ParserRuleContext {
		public TerminalNode NULL() { return getToken(JSONParser.NULL, 0); }
		public NilContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nil; }
	}

	public final NilContext nil() throws RecognitionException {
		NilContext _localctx = new NilContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_nil);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			match(NULL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\17M\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\4"+
		"\f\t\f\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\5\3!\n\3\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\5\4)\n\4\3\5\3\5\3\5\3\5\3\5\5\5\60\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\5\7<\n\7\3\b\3\b\3\b\3\b\3\b\5\bC\n\b\3\t\3\t\3\n\3\n\3\13"+
		"\3\13\3\f\3\f\3\f\2\2\r\2\4\6\b\n\f\16\20\22\24\26\2\3\3\2\t\n\2J\2\30"+
		"\3\2\2\2\4 \3\2\2\2\6(\3\2\2\2\b/\3\2\2\2\n\61\3\2\2\2\f;\3\2\2\2\16B"+
		"\3\2\2\2\20D\3\2\2\2\22F\3\2\2\2\24H\3\2\2\2\26J\3\2\2\2\30\31\5\4\3\2"+
		"\31\3\3\2\2\2\32!\5\6\4\2\33!\5\f\7\2\34!\5\22\n\2\35!\5\20\t\2\36!\5"+
		"\24\13\2\37!\5\26\f\2 \32\3\2\2\2 \33\3\2\2\2 \34\3\2\2\2 \35\3\2\2\2"+
		" \36\3\2\2\2 \37\3\2\2\2!\5\3\2\2\2\"#\7\3\2\2#)\7\4\2\2$%\7\3\2\2%&\5"+
		"\b\5\2&\'\7\4\2\2\')\3\2\2\2(\"\3\2\2\2($\3\2\2\2)\7\3\2\2\2*\60\5\n\6"+
		"\2+,\5\n\6\2,-\7\7\2\2-.\5\n\6\2.\60\3\2\2\2/*\3\2\2\2/+\3\2\2\2\60\t"+
		"\3\2\2\2\61\62\5\20\t\2\62\63\7\b\2\2\63\64\5\4\3\2\64\13\3\2\2\2\65\66"+
		"\7\5\2\2\66<\7\6\2\2\678\7\5\2\289\5\16\b\29:\7\6\2\2:<\3\2\2\2;\65\3"+
		"\2\2\2;\67\3\2\2\2<\r\3\2\2\2=C\5\4\3\2>?\5\4\3\2?@\7\7\2\2@A\5\4\3\2"+
		"AC\3\2\2\2B=\3\2\2\2B>\3\2\2\2C\17\3\2\2\2DE\7\16\2\2E\21\3\2\2\2FG\7"+
		"\17\2\2G\23\3\2\2\2HI\t\2\2\2I\25\3\2\2\2JK\7\13\2\2K\27\3\2\2\2\7 (/"+
		";B";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}