grammar JSON;

json: value;
value: object | array | string | int | TRUE | FALSE | NULL;
object: LEFT_BRACKET RIGHT_BRACKET | LEFT_BRACKET RIGHT_BRACKET;
members: member | member COMMA member;
member: string COLORN value;
array:  LEFT_SQUARE_BRACKET RIGHT_SQUARE_BRACKET | LEFT_SQUARE_BRACKET elements RIGHT_SQUARE_BRACKET;
elements: value | value COMMA value; 
string: DOUBLE_QUOTE LETTER DOUBLE_QUOTE;
int: NUMBER | MINUS NUMBER;

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
NUMBER: [0-9]+;
MINUS: '-'; 
LETTER: [a-Z]*;
