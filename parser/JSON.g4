grammar JSON;

json: value;
value: object | array | num | str | bool | nil;
object: LEFT_BRACKET RIGHT_BRACKET | LEFT_BRACKET members RIGHT_BRACKET;
members: member | member COMMA members;
member: str COLORN value;
array:  LEFT_SQUARE_BRACKET RIGHT_SQUARE_BRACKET | LEFT_SQUARE_BRACKET elements RIGHT_SQUARE_BRACKET;
elements: value | value COMMA elements; 
str: STRING;
num: INTEGER;
bool: TRUE | FALSE;
nil: NULL;

LEFT_BRACKET: '{';
RIGHT_BRACKET: '}';
LEFT_SQUARE_BRACKET: '[';
RIGHT_SQUARE_BRACKET: ']';
COMMA: ',';
COLORN: ':';
TRUE: 'true';
FALSE: 'false';
NULL: 'null';
DOUBLE_QUOTE: '"';
WHITE_SPACE : [ \r\t\n]+ -> skip;
STRING: '"' [a-zA-Z0-9]* '"';
fragment DIGIT: '0'..'9';
INTEGER: '0' | '1'..'9' DIGIT* | '-' '1'..'9' DIGIT*; 
