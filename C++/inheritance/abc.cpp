#include <iostream>
#include <vector>
/* This is what we refer to as a pure virtual class i.e. Abstract Base Class.
   You cannot instantiate this class, only classes derived from this class.
	 In Java, this is known as an Interface. So pls stop making fun of Java :(
	 *** NOTE ***
	 saying a function virtual returntype funcname() = 0; is a contract. You're
	 telling the compiler that you're going to implement that function later in
	 a derived class. If you don't, you'll get a slew of warnings.
*/
class Food {
public:
	virtual int deliciousness() = 0;
	virtual bool isCooked() = 0;
	virtual bool isHealthy() = 0;
	virtual std::string getName() = 0;
};

class Apple : public Food {
public:
	Apple() : name("Delicious") {}
	int deliciousness() { return 7; }
	bool isCooked() {return false; }
	bool isHealthy() { return true;}
	std::string getName() {return name; }
private:
	std::string name;
};

class Cookie : public Food {
public:
	Cookie() : name("Dessert") {}
	int deliciousness() {return 11; }
	bool isCooked() {return true; }
	bool isHealthy() {return false;}
	std::string getName() { return name; }
private:
	std::string name;
};
int main() {
	std::vector<Food*> FOOD;
	FOOD.reserve(5);
	for (int i = 0; i < 5; i++) {
		if (i%2) FOOD.emplace_back(new Apple);
		else FOOD.emplace_back(new Cookie);
	}
	for (Food *food : FOOD)
		std::cout << food->getName() << std::endl;
}
