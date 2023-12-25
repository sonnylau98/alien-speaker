#include<fstream>
#include<iostream>
#include <cstdlib>

using std::string;
using std::cout;
using std::ofstream;
using std::ifstream;
using std::cin;
using std::endl;

unsigned int key=0;
string EN(string en,unsigned int k);
string DE(string de,unsigned int k);

char dec2hexChar(short int n)
{
    if ( 0 <= n && n <= 9 )
    {
        return char( short('0') + n );
    }
	else if ( 10 <= n && n <= 15 )
	{
        return char( short('A') + n - 10 );
    }
	else
	{
        return char(0);
    }
}

short int hexChar2dec(char c)
{
    if ( '0'<=c && c<='9' )
	{
        return short(c-'0');
    }
	else if ( 'a'<=c && c<='f' )
	{
        return ( short(c-'a') + 10 );
    }
	else if ( 'A'<=c && c<='F' )
	{
        return ( short(c-'A') + 10 );
    }
	else
    {
        return -1;
    }
}

string escapeURL(const string &URL)
{
    string result = "";
    for ( unsigned int i=0; i<URL.size(); i++ )
	{
        char c = URL[i];
        if (( '0'<=c && c<='9' ) ||( 'a'<=c && c<='z' ) ||( 'A'<=c && c<='Z' ) ||c=='/' || c=='.')
	    {
            result += c;
        }
	    else
	    {
            int j = (short int)c;
            if ( j < 0 )
		    {
                j += 256;
            }
            int i1, i0;
            i1 = j / 16;
            i0 = j - i1*16;
            result += '%';
            result += dec2hexChar(i1);
            result += dec2hexChar(i0);
        }
    }
    return result;
}

string deescapeURL(const string &URL)
{
    string result = "";
    for ( unsigned int i=0; i<URL.size(); i++ )
	{
        char c = URL[i];
        if ( c != '%' )
		{
            result += c;
        }
		else
		{
            char c1 = URL[++i];
            char c0 = URL[++i];
            int num = 0;
            num += hexChar2dec(c1) * 16 + hexChar2dec(c0);
            result += char(num);
        }
    }
    return result;
}

int main()
{
	void fun1();
	void fun2();
	cout<<"Hello, there is cryption program!\n";
	cout<<"\n Input the key:";
	if(scanf("%u",&key)==0)
    {
        cout<<"\n invalid input\n";
        system("pause");
        exit(1);
    }
	int num;
	while(1)
	{
		cout<<"\n Message[1], Decoding[2], Change key[3], Exit[4].\n Please input:";

		if(scanf("%d",&num)==0)
		{
			cout<<"\n invalid input\n";
			system("pause");
			exit(1);
		 }
		switch(num)
		{
			case 1 : fun1();
			    break;
			case 2 : fun2();
			    break;
			case 3 :
			{
			    cout<<"\n Input the key:";
	            if(scanf("%u",&key)==0)
		        {
                    cout<<"\n invalid input\n";
                    system("pause");
			        exit(1);
		        }
			}
			    break;
			case 4 :
			{
				cout<<"\n Bye!\n\n";
				system("pause");
				exit(0);
			}
			    break;
			default : cout<<"\n Warning: please input correctly.\n";
		 }
	}

	return 0;
 }

void fun1()
{
	string Text="";
	string cryptograph="";
	string mi="";
	cout<<"\n Input message:\n";
    cin.ignore();
	getline(cin,Text);
	cryptograph=escapeURL(Text);
	mi=EN(cryptograph,key);
    ofstream outfile;
    outfile.open("Ciphertext.txt");
    outfile << mi << endl;
    cout<<"\n Done! Check the file Ciphertext.txt please.\n";
    outfile.close();
}

void fun2()
{
	string Text="";
	string cryptograph="";
	string ming="";
	cout<<"\n Put the ciphertext in Ciphertext.txt.\n After confirmation";
	system("pause");
	ifstream infile;
    infile.open("Ciphertext.txt");
    infile >>cryptograph;
    ming=DE(cryptograph,key);
    Text=deescapeURL(ming);
	cout<<"\n The message actually is: \n"<<Text<<endl;
    infile.close();
}

string EN(string en,unsigned int k) //URL
{
	string E;
	int i,j;
	for(i=0;i<en.size();)
	{
		for(j=0;(i<en.size())&&(j<32);i++,j++)
		{
			if(i==0)
			{
				j=k%32;
			}
			E.push_back(en[i]+j);
		}
	}
	return E;
}

string DE(string de,unsigned int k)
{
	string D;
	int i,j;
	for(i=0;i<de.size();)
	{
		for(j=0;(i<de.size())&&(j<32);i++,j++)
		{
			if(i==0)
			{
				j=k%32;
			}
			D.push_back(de[i]-j);
		}
	}
	return D;
}
