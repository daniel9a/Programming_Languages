/*******************************************/
/**    Your solution goes in this file    **/ 
/*******************************************/

/* Part 1*/

year_1953_1996_novels(X) :- /*novels made in 1996 or 1953 */
    novel(X,1996) ; novel(X,1953).

    
period_1800_1900_novels(X) :- novel(X,Y), /*novels made in between the years 1800 and 1900 */
    Y >= 1800, Y =< 1900.

/*people who are a fan of lord of the rings*/    
lotr_fans(X) :- 
    fan(X, Y), 
    member(the_lord_of_the_rings, Y).

/*see if a is a part of list f and vice versa*/
compare(A, F) :-
    member(C, F), 
    member(C, A),
    !.  

/*authors chandler is a fan of*/ 
author_names(X) :- 
    fan(chandler, Ch),
    author(X, Author), 
    compare(Ch, Author).
    
fans_names(X) :-
    author(brandon_sanderson, BS),
    fan(X, Fan),
    compare(Fan, BS).

/*finds novels that are common between pheobe ross and monica */
mutual_novels(X) :- /*check phoebe and ross */
    fan(phoebe, P),
    fan(ross, R),
    nCheck(X, P, R).
    
mutual_novels(X) :-
    fan(phoebe, P),
    fan(monica, M),
    nCheck(X, P, M).
    
mutual_novels(X) :-
    fan(ross, R),
    fan(monica, M),
    nCheck(X, R, M).
    
nCheck(X, [H|T], L2) :-
    X = H,
    member(H,L2).
    
nCheck(X, [_|T], L2) :-
    nCheck(X, T, L2).    
    


/* PART 2 */

isMember(X, [X|_]).

isMember(X, [_|T]) :-
    isMember(X, T).
   
isUnion([], L, L).

isUnion([H|T1], L2, L3) :-
    isMember(H, L2),
    !,
    isUnion(T1, L2, L3).
    
isUnion([H|T1], L2, [H|T3]) :-
    isUnion(T1, L2, T3).
    
    
isIntersection([], _, []).

isIntersection([H1|T1], L2, [H1|L3]) :-
    isMember(H1, L2), 
    !,
    isIntersection(T1, L2, L3).

isIntersection([_|T1], L2, L3) :-
    isIntersection(T1, L2, L3).


isEqual(L1, L2) :-
    isIntersection(L1, L2, L3),
    L3 = L1.
 

powerSet([],[[]]).
powerSet([H|T],L)  :-
	powerSet(T,L1),
	appendEach([H],L1,L2),
	myappend(L1,L2,L).

appendEach(X,[H|[]],[H1]):-
	myappend(X,H,H1).
appendEach(X,[H|T],[H1|T1]):-
	myappend(X,H,H1),
	appendEach(X,T,T1).
	
myappend([],L,L).
myappend([X|L1],L2,[X|L3]) :-
    myappend(L1,L2,L3).


/*Part 3*/

unsafe(state(X, Y, Y, C)) :- opposite(X,Y).
unsafe(state(X, W, Y, Y)) :- opposite(X,Y).

opposite(right,left). 
opposite(left,right). 

arc(X,A,B).
arc(take(wolf,X,Y), state(X,X,G,C),state(Y,Y,G,C)) :- 
    opposite(X,Y), 
    \+(unsafe(state(Y,Y,G,C))).

arc(take(goat,X,Y), state(X,W,X,C),state(Y,W,Y,C)) :- 
    opposite(X,Y), 
    \+(unsafe(state(Y,W,Y,C))).

arc(take(cabbage,X,Y), state(X,W,G,X),state(Y,W,G,Y)) :- 
    opposite(X,Y), 
    \+(unsafe(state(Y,W,Y,C))).
    
arc(take(none,X,Y), state(X,W,G,C),state(Y,W,G,C)) :- 
    opposite(X,Y), 
    \+(unsafe(state(Y,W,G,C))).



output(S1,S2,State0,State1,State2,State3,State4,State5,State6,State7,State8,State9,State10,State11,State12,State13,State14,State15):-
    (
      (S1=State0,S2=State9) 
        -> write(take(cabbage,left,right)),nl; 
      ((S1=State0,S2=State10) 
        -> write(take(goat,left,right)),nl;
      ((S1=State10,S2=State2) 
        -> write(take(none,right,left)),nl;
      ((S1=State2,S2=State14) 
        -> write(take(wolf,left,right)),nl;
      ((S1=State14,S2=State4) 
        -> write(take(goat,right,left)),nl; 
      ((S1=State4,S2=State13) 
        -> write(take(cabbage,left,right)),nl;
      ((S1=State13,S2=State5) 
        -> write(take(none,right,left)),nl;
      ((S1=State5,S2=State15) 
        -> write(take(goat,left,right)),nl; 
      ((S1=State2,S2=State11) 
        -> write(take(cabbage,left,right)),nl; 
      ((S1=State11,S2=State1) 
        -> write(take(goat,right,left)),nl; 
      true)))))))))
    ),!.



check([]).
check([H|[]]).
check([State0|T]) :-
    [State1|T1] = T,
    output(State0,State1,state(left,left,left,left),
    state(left,left,left,right),state(left,left,right,left),
    state(left,left,right,right),state(left,right,left,left),
    state(left,right,left,right),state(left,right,right,left),
    state(left,right,right,right),state(right,left,left,left),
    state(right,left,left,right),state(right,left,right,left),
    state(right,left,right,right),state(right,right,left,left),
    state(right,right,left,right),state(right,right,right,left),
    state(right,right,right,right)),
    check(T),!.


reverse([],[]).
reverse([H|T],L1) :-
  reverse(T,L2), myappend(L2,[H],L1).


path(List, Goal, Goal):-
    Y=[Goal|List],
    reverse(Y,X),
    check(X).

path(List, State, Goal):-
    \+(unsafe(State)),
    \+(unsafe(Goal)),
	arc(A,State, NextState),
	\+(isMember(NextState, List)),
	path([NextState|List], NextState, Goal),!.
	

solve(F1,W1,G1,C1,F2,W2,G2,C2) :-
    go(state(F1,W1,G1,C1), state(F2,W2,G2,C2)).

go(A,B) :-
    S = [A],
    path(S, A, B).


