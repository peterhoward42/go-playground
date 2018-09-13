# Observations on testify

I perceive two different uses of mock objects. 

Firstly alternatives to real objects on which the code under test is dependent
on, and which are either easier to bring up than the real ones, or can be made
more deterministic to make testing easier.

Secondly to be provided as substitutes for real objects that the code under test
is dependent on, which provide sham implementations, but crucially record their
call history, for subsequent scrutiny.

From a first look testify's mocks do the second thing very well, but don't seem
to make it any easier to do the first thing than coding a trivial implementation
of the interface required.

