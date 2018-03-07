#include <iostream>

// example of how to reuse code from the parent class in the derived class 
class Person {
public:
	Person() : name("Spudley"), year(2) { std::cout << "Person constructed!\n";}
	Person(std::string n, int y) : name(n), year(y) { std::cout << "Person conversion!\n"; }
	friend std::ostream& operator<<(std::ostream& os, const Person &p)
	{return os << p.name << " " << p.year;}
private:
std::string name;
int year;
};

class Student : public Person {
public:
	Student() {std::cout << "Student constructed!\n";}
	Student(std::string n, int y) : Person(n, y) { std::cout << "Student Conversion!\n";}
private:
};

int main() {
	Person p1;
	std::cout << "p1 created!" << p1 << std::endl;
	Person p2("Potato", 22);
	std::cout << "p2 created!\n" << p2 << std::endl;
	Student s1;
	std::cout << "s1 created!\n" << s1 << std::endl;
	Student s2("Tater", 12);
	std::cout << "s2 created!\n" << s2 << std::endl;
}
